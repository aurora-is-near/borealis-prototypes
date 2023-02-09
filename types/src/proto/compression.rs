use std::io::{self, Read};
use thiserror::Error;
use zstd::stream;

#[derive(Debug, Error)]
pub enum DecodeError {
    #[error("{0}")]
    Io(#[from] io::Error),
    #[error("{0}")]
    Prost(#[from] prost::DecodeError),
}

pub trait CompressedMessage: prost::Message {
    /// Encodes and compresses the message to a newly allocated buffer.
    ///
    /// Result will be in the zstd frame format. A `compression_level` of 0 uses zstd's default (currently 3).
    fn encode_compressed(&self, compression_level: i32) -> Result<Vec<u8>, io::Error>;

    /// Decompresses and decodes an instance of the message from a buffer.
    ///
    /// The entire buffer will be consumed. The input data must be in the zstd frame format.
    fn decode_compressed(source: impl Read) -> Result<Self, DecodeError>
    where
        Self: Sized;
}

impl<T: prost::Message + Default> CompressedMessage for T {
    fn encode_compressed(&self, compression_level: i32) -> Result<Vec<u8>, io::Error> {
        let source = self.encode_to_vec();

        stream::encode_all(&source[..], compression_level)
    }

    fn decode_compressed(source: impl Read) -> Result<Self, DecodeError> {
        let bytes = stream::decode_all(source)?;

        Ok(Self::decode(&bytes[..])?)
    }
}
