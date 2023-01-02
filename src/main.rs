mod conversions;
mod proto;

use borealis_types::message::Message;
use borealis_types::payloads::NEARBlock;
use chrono::prelude::*;
use clap::Parser;
use nats::jetstream::{ConsumerConfig, PullSubscribeOptions};
use nats::ServerAddress;
use std::error::Error;
use std::path::Path;
use std::str::FromStr;

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
    pub async fn convert(
        path: impl AsRef<Path>,
        urls: impl AsRef<str>,
        subject: impl AsRef<str>,
    ) -> Result<(), Box<dyn Error>> {
        let connection = nats::Options::with_credentials(path).connect(urls.as_ref())?;

        println!("Connected");

        let stream = nats::jetstream::new(connection);
        let mut cons = ConsumerConfig::default();
        let options = PullSubscribeOptions::new().consumer_config(cons);
        let sub = stream.pull_subscribe_with_options(subject.as_ref(), &options).unwrap();

        sub.fetch_with_handler(1, |message| {
            let parsed = Message::<NEARBlock>::from_cbor(&message.data).unwrap();
            println!("{:#?}", parsed);

            let protobuf = proto::Messages::from(parsed);
            println!("{:#?}", protobuf);

            Ok(())
        })?;

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
