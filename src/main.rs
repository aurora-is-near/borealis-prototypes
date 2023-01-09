mod conversions;
mod proto;

use borealis_types::payloads::NEARBlock;
use chrono::prelude::*;
use clap::Parser;
use nats::jetstream::{ConsumerConfig, DeliverPolicy, PullSubscribeOptions};
use nats::ServerAddress;
use std::error::Error;
use std::path::Path;
use std::str::FromStr;
use prost::Message;
use time::macros::datetime;

#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
struct Config {
    #[arg(short, long)]
    creds: String,
    #[arg(short, long)]
    nats_server: String,
    #[arg(short = 's', long)]
    nats_subject: String,
}

struct Convertor;

impl Convertor {
    const START_SEQ_2022_01: u64 = 21557350;
    const START_SEQ_2022_02: u64 = Self::START_SEQ_2022_01 + 2120297;
    const START_SEQ_2022_05: u64 = Self::START_SEQ_2022_01 + 8189072;
    const START_SEQ_2022_06: u64 = Self::START_SEQ_2022_01 + 10323540;
    const START_SEQ_2022_07: u64 = Self::START_SEQ_2022_01 + 12407004;

    pub async fn convert(
        path: impl AsRef<Path>,
        urls: impl AsRef<str>,
        subject: impl AsRef<str>,
    ) -> Result<(), Box<dyn Error>> {
        let connection = nats::Options::with_credentials(path).connect(urls.as_ref())?;
        let stream = nats::jetstream::new(connection);
        let mut cons = ConsumerConfig::default();
        cons.opt_start_seq = Some(Self::START_SEQ_2022_07);
        cons.deliver_policy = DeliverPolicy::ByStartSeq;
        let mut seq = cons.opt_start_seq.unwrap_or_default();
        let options = PullSubscribeOptions::new().consumer_config(cons);
        let sub = stream.pull_subscribe_with_options(subject.as_ref(), &options).unwrap();
        println!("height;seq;date;cbor;proto;diff");

        for _ in (Self::START_SEQ_2022_01 / 100)..(Self::START_SEQ_2022_02 / 100) {
            sub.fetch_with_handler(100, |message| {
                let old_size = message.data.len();
                let parsed = borealis_types::message::Message::<NEARBlock>::from_cbor(&message.data).unwrap();
                let date = chrono::NaiveDateTime::from_timestamp_millis(
                    (parsed.payload.block.header.timestamp as f64 * 0.000001) as i64
                ).unwrap();
                let id = parsed.payload.block.header.height;
                seq += 1;

                let protobuf = proto::Messages::from(parsed);
                let new_size = protobuf.into_inner()
                    .into_iter()
                    .map(|v: proto::Message| v.encoded_len())
                    .fold(0, |acc, v| acc + v);
                println!("{};{};\"{}\";{};{};{}", id, seq, date, old_size, new_size, new_size as f64 / old_size as f64);

                Ok(())
            })?;
        }

        Ok(())
    }
}

#[tokio::main]
async fn main() {
    let config = Config::parse();

    Convertor::convert(config.creds, config.nats_server, config.nats_subject)
        .await
        .unwrap();
}
