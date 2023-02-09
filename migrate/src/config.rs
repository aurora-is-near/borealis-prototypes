use clap::Parser;
use std::num::ParseIntError;
use std::str::FromStr;
use thiserror::Error;

#[derive(Debug, clap::Subcommand)]
pub(crate) enum Command {
    /// Runs on a range from `sequence_start` to `sequence_stop` and generates v3 NEAR blocks into `nats_output_server`
    /// from v2 NEAR blocks on `nats_server`.
    Migrate,
    /// Runs on a range from `sequence_start` to `sequence_stop` and generates statistics concerning size difference
    /// in conversion from v2 to v3 NEAR blocks.
    Convert,
    /// Prints a block to stdout and then quits.
    Print,
}

#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
pub(crate) struct Config {
    /// The maximum number of messages that can be pending at once for acknowledgement when processing.
    #[arg(short = 'i', long, default_value = "10")]
    pub(crate) batch_size: usize,
    /// Used for compression of the v3 NEAR blocks ranging from -7 (fastest) to 22 (best compression ratio).
    #[arg(short = 'l', long, default_value = "9")]
    pub(crate) compression_level: i16,
    /// Dumps processed blocks into files.
    #[arg(short = 'f', long)]
    pub(crate) dump_blocks: bool,
    /// Credentials for the `nats_server`.
    #[arg(short, long)]
    pub(crate) nats_creds: String,
    /// The NATS server address to read the input messages from.
    #[arg(short = 'a', long)]
    pub(crate) nats_server: String,
    /// The subject of the messages to process.
    #[arg(short = 's', long)]
    pub(crate) nats_subject: String,
    /// Credentials for the `nats_output_server`.
    #[arg(short = 'k', long)]
    pub(crate) nats_output_creds: Option<String>,
    /// The NATS server address to publish the output messages to. The payloads are protobuf encoded and zstd compressed
    /// NEAR blocks split by header and shards. Each is published on a different subject.
    #[arg(short = 'u', long)]
    pub(crate) nats_output_server: String,
    /// The subject of the published NATS messages that contain NEAR block header as a payload.
    ///
    /// The payload is encoded in protobuf and compressed using zstd.
    #[arg(short = 'j', long)]
    pub(crate) nats_output_subject_header: String,
    /// Prefix for the published NATS message subject. The prefix is then followed by a shard ID, expected as a
    /// numerical index starting from 0 and used as a subject for the published message that contains NEAR block shard
    /// data as a payload.
    ///
    /// The payload is encoded in protobuf and compressed using zstd.
    #[arg(short = 'd', long)]
    pub(crate) nats_output_subject_shards: String,
    /// Prints CSV compatible header for the CSV compatible values printed in the log output.
    #[arg(short = 'e', long, default_value = "false")]
    pub(crate) print_header: bool,
    /// First sequence number on the source stream to start publishing on. It is optional, if not provided, it runs from
    /// number 1 which is the lowest.
    #[arg(short = 'r', long, default_value = "1")]
    pub(crate) sequence_start: u64,
    /// Final sequence number on the source stream to stop publishing on. Program runs until this number is hit. It is
    /// optional, if not provided then it will run until interrupted.
    #[arg(short = 'o', long)]
    pub(crate) sequence_stop: Option<u64>,
    /// Suppress `Nats-Expected-Last-Msg-Id` NATS message header when publishing first message.
    /// Only effective when `sequence_start` is higher than `1`, otherwise the `Nats-Expected-Last-Msg-Id` cannot be
    /// included as there is no previous message.
    #[arg(short = 'x', long)]
    pub(crate) skip_expect: bool,
    /// Maximum `timeout` duration for the connection to the incoming stream. If a message is not received within this
    /// `timeout` duration, the connection is reestablished.
    ///
    /// Example: 10s, 500ms
    #[arg(short = 't', long, default_value = "10s")]
    pub(crate) timeout: Duration,
    #[command(subcommand)]
    pub(crate) command: Command,
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
