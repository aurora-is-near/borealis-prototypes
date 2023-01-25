use anyhow::anyhow;
use async_nats::jetstream::consumer::{Consumer, DeliverPolicy};
use async_nats::jetstream::{consumer, Context};
use async_nats::ConnectOptions;
use borealis_proto_types as proto;
use borealis_types::payloads::NEARBlock;
use futures::{stream::FuturesUnordered, StreamExt};
use futures_util::TryStreamExt;
use prost::Message;
use std::io::Write;
use std::path::Path;

pub(crate) struct Convertor;

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
        sequence_stop: Option<u64>,
        print_header: bool,
        batch_size: usize,
        timeout: std::time::Duration,
        compression_level: i32,
    ) -> anyhow::Result<()> {
        if print_header {
            log::info!("seq;height;date;cbor;proto;proto_zstd;diff;diff_zstd");
        }

        let shards = 4;
        let mut seq = sequence_start;

        loop {
            let jetstream = Self::jetstream(output_creds.as_ref(), &output_urls).await?;
            let consumer = Self::consumer(&path, &urls, &subject, seq).await?;
            let mut sequence = consumer.sequence(batch_size).map_err(|e| anyhow!(e))?;

            'sequence: loop {
                match sequence_stop {
                    Some(sequence_stop) if seq >= sequence_stop => return Ok(()),
                    _ => {}
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
                                let last_shard_id = protobuf.len() - 1;

                                let mut futures = FuturesUnordered::new();
                                let empty_shards = (last_shard_id..shards).map(|shard_id| {
                                    jetstream.publish(
                                        format!("{}{}", subject_shard.as_ref(), shard_id),
                                        Default::default()
                                    )
                                });
                                protobuf.into_iter().filter_map(|msg| {
                                    let source = msg.encode_to_vec();
                                    let mut compressed: Vec<u8> = Vec::new();
                                    zstd::stream::copy_encode(&source[..], &mut compressed, compression_level).unwrap();

                                    new_size += source.len();
                                    compressed_size += compressed.len();

                                    Some(jetstream.publish(match msg.payload.expect("Payload is mandatory") {
                                        proto::message::Payload::NearBlockHeader(..) => subject_header.as_ref().to_string(),
                                        proto::message::Payload::NearBlockShard(shard) => format!("{}{}", subject_shard.as_ref(), shard.shard_id),
                                        _ => return None,
                                    }, compressed.into()))
                                }).chain(empty_shards).for_each(|publish_fut| futures.push(publish_fut));

                                let mut has_error = false;
                                while let Some(result) = futures.next().await {
                                    if let Err(error) = result {
                                        log::error!("{error:?}");
                                        has_error = true;
                                    }
                                }
                                if has_error {
                                    return Err(anyhow!("Error occurred while publishing messages."));
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
        sequence_stop: Option<u64>,
        print_header: bool,
        batch_size: usize,
        timeout: std::time::Duration,
        compression_level: i32,
    ) -> anyhow::Result<()> {
        if print_header {
            println!("seq;height;date;cbor;proto;proto_zstd;diff;diff_zstd");
        }

        let mut seq = sequence_start;

        loop {
            let consumer = Self::consumer(&path, &urls, &subject, seq).await?;
            let mut sequence = consumer.sequence(batch_size).unwrap();

            'sequence: loop {
                match sequence_stop {
                    Some(sequence_stop) if seq >= sequence_stop => return Ok(()),
                    _ => {}
                }

                'batch: while let Some(mut batch) = sequence.try_next().await.unwrap() {
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
