mod app;
mod config;

use crate::app::Convertor;
use crate::config::{Command, Config};
use clap::Parser;

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
