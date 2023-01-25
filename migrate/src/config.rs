use clap::Parser;
use std::num::ParseIntError;
use std::str::FromStr;
use thiserror::Error;

#[derive(Debug, clap::Subcommand)]
pub(crate) enum Command {
    Migrate,
    Convert,
    Print,
}

#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
pub(crate) struct Config {
    #[arg(short = 'i', long, default_value = "10")]
    pub(crate) batch_size: usize,
    #[arg(short = 'l', long, default_value = "9")]
    pub(crate) compression_level: u16,
    #[arg(short, long)]
    pub(crate) nats_creds: String,
    #[arg(short = 'a', long)]
    pub(crate) nats_server: String,
    #[arg(short = 's', long)]
    pub(crate) nats_subject: String,
    #[arg(short = 'k', long)]
    pub(crate) nats_output_creds: Option<String>,
    #[arg(short = 'u', long)]
    pub(crate) nats_output_server: String,
    #[arg(short = 'j', long)]
    pub(crate) nats_output_subject_header: String,
    #[arg(short = 'd', long)]
    pub(crate) nats_output_subject_shards: String,
    #[arg(short = 'e', long, default_value = "false")]
    pub(crate) print_header: bool,
    #[arg(short = 'r', long, default_value = "1")]
    pub(crate) sequence_start: u64,
    #[arg(short = 'o', long)]
    pub(crate) sequence_stop: Option<u64>,
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
