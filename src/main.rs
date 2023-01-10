mod conversions;
mod proto;

use anyhow::anyhow;
use async_nats::jetstream::consumer::{Consumer, DeliverPolicy};
use async_nats::jetstream::{consumer, Context};
use async_nats::ConnectOptions;
use borealis_types::payloads::NEARBlock;
use clap::Parser;
use futures_util::TryStreamExt;
use prost::Message;
use std::io::Write;
use std::num::ParseIntError;
use std::path::Path;
use std::str::FromStr;
use thiserror::Error;

#[derive(Debug, clap::Subcommand)]
enum Command {
    Migrate,
    Convert,
    Print,
}

#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
pub(crate) struct Config {
    #[arg(short, long)]
    creds: String,
    #[arg(short = 'a', long)]
    nats_server: String,
    #[arg(short = 's', long)]
    nats_subject: String,
    #[arg(short = 'r', long)]
    sequence_start: u64,
    #[arg(short = 'o', long)]
    sequence_stop: u64,
    #[arg(short = 'e', long, default_value = "false")]
    print_header: bool,
    #[arg(short = 'i', long, default_value = "10")]
    batch_size: usize,
    #[arg(short = 't', long, default_value = "1s")]
    timeout: Duration,
    #[arg(short = 'l', long, default_value = "9")]
    compression_level: u16,
    #[arg(short = 'k', long)]
    nats_output_creds: Option<String>,
    #[arg(short = 'u', long)]
    nats_output_server: String,
    #[arg(short = 'j', long)]
    nats_output_subject_header: String,
    #[arg(short = 'd', long)]
    nats_output_subject_shards: String,
    #[command(subcommand)]
    command: Command,
}

#[derive(Debug, Error)]
pub(crate) enum DurationParseError {
    #[error("Unknown duration format")]
    UnknownFormat,
    #[error("Invalid duration number")]
    NumberFormat(#[from] ParseIntError),
}

#[derive(Debug, Clone)]
pub(crate) struct Duration(pub(crate) std::time::Duration);

impl From<Duration> for std::time::Duration {
    fn from(value: Duration) -> Self {
        value.0
    }
}

impl FromStr for Duration {
    type Err = DurationParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        Ok(Self(if let Some(number) = s.strip_suffix("ms") {
            let millis = u64::from_str(number)?;

            std::time::Duration::from_millis(millis)
        } else if let Some(number) = s.strip_suffix("us") {
            let micros = u64::from_str(number)?;

            std::time::Duration::from_micros(micros)
        } else if let Some(number) = s.strip_suffix("ns") {
            let nanos = u64::from_str(number)?;

            std::time::Duration::from_nanos(nanos)
        } else if let Some(number) = s.strip_suffix('s') {
            let secs = u64::from_str(number)?;

            std::time::Duration::from_secs(secs)
        } else {
            return Err(DurationParseError::UnknownFormat);
        }))
    }
}

struct Convertor;

impl Convertor {
    pub async fn migrate(
        path: impl AsRef<Path>,
        urls: impl AsRef<str>,
        subject: impl AsRef<str>,
        output_creds: Option<impl AsRef<Path>>,
        output_urls: impl AsRef<str>,
        subject_header: impl AsRef<str>,
        subject_shard: impl AsRef<str>,
        sequence_start: u64,
        sequence_stop: u64,
        print_header: bool,
        batch_size: usize,
        timeout: std::time::Duration,
        compression_level: i32,
    ) -> anyhow::Result<()> {
        if print_header {
            log::info!("seq;height;date;cbor;proto;proto_zstd;diff;diff_zstd");
        }

        let mut seq = sequence_start;

        loop {
            let jetstream = Self::jetstream(output_creds.as_ref(), &output_urls).await?;
            let consumer = Self::consumer(&path, &urls, &subject, seq).await?;
            let mut sequence = consumer.sequence(batch_size).map_err(|e| anyhow!(e))?;

            'sequence: loop {
                if seq >= sequence_stop {
                    return Ok(());
                }

                match sequence.try_next().await.map_err(|e| anyhow!(e))? {
                    None => return Ok(()),
                    Some(mut batch) => 'batch: loop {
                        let sleep = tokio::time::sleep(timeout);
                        tokio::pin!(sleep);
                        tokio::select! {
                            _ = &mut sleep => {
                                log::error!("Timeout with server. Restarting connection.");
                                break 'sequence;
                            },
                            message = batch.try_next() => if let Ok(Some(message)) = message {
                                let old_size = message.message.payload.len();
                                let parsed = borealis_types::message::Message::<NEARBlock>::from_cbor(&message.message.payload)?;

                                message.ack().await.map_err(|e| anyhow!(e))?;

                                let id = parsed.payload.block.header.height;
                                let date = chrono::NaiveDateTime::from_timestamp_millis(
                                    (parsed.payload.block.header.timestamp as f64 * 0.000001) as i64
                                ).unwrap_or_default();
                                let protobuf = proto::Messages::from(parsed).into_inner();
                                let mut new_size = 0;
                                let mut compressed_size = 0;

                                for msg in protobuf {
                                    let source = msg.encode_to_vec();
                                    let mut compressed: Vec<u8> = Vec::new();
                                    zstd::stream::copy_encode(&source[..], &mut compressed, compression_level)?;

                                    new_size += source.len();
                                    compressed_size += compressed.len();

                                    jetstream.publish(match msg.payload.expect("Payload is mandatory") {
                                        proto::message::Payload::NearBlockHeader(..) => subject_header.as_ref().to_string(),
                                        proto::message::Payload::NearBlockShard(shard) => format!("{}{}", subject_shard.as_ref(), shard.shard_id),
                                        _ => continue,
                                    }, compressed.into()).await.map_err(|e| anyhow!(e))?.await.map_err(|e| anyhow!(e))?;
                                }

                                log::info!("{seq};{id};\"{date}\";{old_size};{new_size};{compressed_size};{};{}",
                                    new_size as f64 / old_size as f64, compressed_size as f64 / old_size as f64
                                );
                                seq += 1;
                            } else {
                                break 'batch;
                            }
                        }
                    },
                }
            }
        }
    }

    pub async fn convert(
        path: impl AsRef<Path>,
        urls: impl AsRef<str>,
        subject: impl AsRef<str>,
        sequence_start: u64,
        sequence_stop: u64,
        print_header: bool,
        batch_size: usize,
        timeout: std::time::Duration,
        compression_level: i32,
    ) -> anyhow::Result<()> {
        if print_header {
            println!("seq;height;date;cbor;proto;proto_zstd;diff;diff_zstd");
        }

        let mut seq = sequence_start;

        while seq < sequence_stop {
            let consumer = Self::consumer(&path, &urls, &subject, seq).await?;
            let mut sequence = consumer.sequence(batch_size).unwrap();

            'sequence: loop {
                let mut batch = sequence.try_next().await.unwrap().unwrap();

                'batch: loop {
                    let sleep = tokio::time::sleep(timeout);
                    tokio::pin!(sleep);
                    tokio::select! {
                        _ = &mut sleep => {
                            log::error!("Timeout with server. Restarting connection.");
                            break 'sequence;
                        },
                        message = batch.try_next() => {
                            if let Ok(Some(message)) = message {
                                log::info!("Message {seq} received.");

                                let old_size = message.message.payload.len();
                                let parsed = borealis_types::message::Message::<NEARBlock>::from_cbor(&message.message.payload).unwrap();
                                let date = chrono::NaiveDateTime::from_timestamp_millis(
                                    (parsed.payload.block.header.timestamp as f64 * 0.000001) as i64
                                ).unwrap();
                                let id = parsed.payload.block.header.height;

                                message.ack().await.unwrap();

                                let protobuf = proto::Messages::from(parsed).into_inner();
                                let new_size: usize = protobuf
                                    .iter()
                                    .map(|v: &proto::Message| v.encoded_len())
                                    .sum();
                                let compressed_size: usize = protobuf
                                    .into_iter()
                                    .map(|v: proto::Message| {
                                        let source = v.encode_to_vec();
                                        let mut compressed: Vec<u8> = Vec::new();
                                        zstd::stream::copy_encode(&source[..], &mut compressed, compression_level).unwrap();
                                        compressed.len()
                                    })
                                    .sum();

                                println!("{seq};{id};\"{date}\";{old_size};{new_size};{compressed_size};{};{}",
                                    new_size as f64 / old_size as f64, compressed_size as f64 / old_size as f64
                                );
                                seq += 1;
                            } else {
                                break 'batch;
                            }
                        }
                    }
                }
            }
        }

        Ok(())
    }

    pub async fn print(
        path: impl AsRef<Path>,
        urls: impl AsRef<str>,
        subject: impl AsRef<str>,
        sequence_id: u64,
    ) -> anyhow::Result<()> {
        'sequence: loop {
            let consumer = Self::consumer(&path, &urls, &subject, sequence_id).await?;
            let mut iter = consumer.sequence(1).unwrap();
            let mut batch = iter.try_next().await.unwrap().unwrap();
            let sleep = tokio::time::sleep(std::time::Duration::from_secs(5));
            tokio::pin!(sleep);
            tokio::select! {
                _ = &mut sleep => {
                    eprintln!("Timeout with server. Restarting connection.");
                    continue 'sequence;
                },
                message = batch.try_next() => {
                    if let Ok(Some(message)) = message {
                        let parsed = borealis_types::message::Message::<NEARBlock>::from_cbor(&message.message.payload).unwrap();
                        println!("{parsed:#?}");

                        let protobuf = proto::Messages::from(parsed).into_inner();
                        println!("{protobuf:#?}");

                        std::io::stdout().flush().unwrap();
                        message.ack().await.unwrap();
                    } else {
                        eprintln!("Error getting message: {message:?}");
                    }
                    break 'sequence;
                }
            }
        }

        Ok(())
    }

    async fn jetstream(path: Option<impl AsRef<Path>>, urls: impl AsRef<str>) -> anyhow::Result<Context> {
        let options = match path {
            Some(path) => ConnectOptions::with_credentials_file(path.as_ref().to_path_buf()).await?,
            None => ConnectOptions::default(),
        };
        let client = async_nats::connect_with_options(urls.as_ref(), options).await?;

        Ok(async_nats::jetstream::new(client))
    }

    async fn consumer(
        path: impl AsRef<Path>,
        urls: impl AsRef<str>,
        subject: impl AsRef<str>,
        start_seq: u64,
    ) -> anyhow::Result<Consumer<consumer::pull::Config>> {
        Self::jetstream(Some(path), urls)
            .await?
            .get_stream("v2_mainnet_near_blocks")
            .await
            .map_err(|e| anyhow!(e))?
            .get_or_create_consumer(
                "borealis_proto",
                consumer::pull::Config {
                    deliver_policy: DeliverPolicy::ByStartSequence {
                        start_sequence: start_seq,
                    },
                    filter_subject: subject.as_ref().to_owned(),
                    ack_policy: consumer::AckPolicy::None,
                    inactive_threshold: std::time::Duration::from_secs(20),
                    ..Default::default()
                },
            )
            .await
            .map_err(|e| anyhow!(e))
    }
}

#[tokio::main]
async fn main() {
    env_logger::init();
    let config = Config::parse();

    match config.command {
        Command::Convert => Convertor::convert(
            config.creds,
            config.nats_server,
            config.nats_subject,
            config.sequence_start,
            config.sequence_stop,
            config.print_header,
            config.batch_size,
            config.timeout.into(),
            config.compression_level as i32,
        )
        .await
        .unwrap(),
        Command::Print => Convertor::print(
            config.creds,
            config.nats_server,
            config.nats_subject,
            config.sequence_start,
        )
        .await
        .unwrap(),
        Command::Migrate => Convertor::migrate(
            config.creds,
            config.nats_server,
            config.nats_subject,
            config.nats_output_creds,
            config.nats_output_server,
            config.nats_output_subject_header,
            config.nats_output_subject_shards,
            config.sequence_start,
            config.sequence_stop,
            config.print_header,
            config.batch_size,
            config.timeout.into(),
            config.compression_level as i32,
        )
        .await
        .unwrap(),
    }
}
