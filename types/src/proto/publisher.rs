use crate::proto;
use async_nats::header::{NATS_EXPECTED_LAST_MESSAGE_ID, NATS_MESSAGE_ID};
use async_nats::HeaderMap;
use async_trait::async_trait;
use bytes::Bytes;
use near_primitives::types::ShardId;
use prost::Message;
use std::error::Error;
use thiserror::Error;

impl From<Box<dyn std::error::Error + Send + Sync>> for PublishError {
    fn from(value: Box<dyn Error + Send + Sync>) -> Self {
        Self {
            message: value.to_string(),
        }
    }
}

#[derive(Clone, Debug, Error, PartialEq)]
#[error("{message}")]
pub struct PublishError {
    message: String,
}

#[async_trait]
pub trait Publisher {
    async fn publish_with_headers(
        &self,
        subject: String,
        headers: HeaderMap,
        payload: Bytes,
    ) -> Result<(), Box<dyn std::error::Error + Send + Sync>>;
}

#[async_trait]
impl Publisher for async_nats::jetstream::Context {
    async fn publish_with_headers(
        &self,
        subject: String,
        headers: HeaderMap,
        payload: Bytes,
    ) -> Result<(), Box<dyn Error + Send + Sync>> {
        self.publish_with_headers(subject, headers, payload).await?.await?;

        Ok(())
    }
}

#[derive(Debug, Clone)]
pub struct EncodingStats {
    pub encoded_size: usize,
    pub compressed_size: usize,
}

impl EncodingStats {
    pub fn new(encoded_size: usize, compressed_size: usize) -> Self {
        Self {
            encoded_size,
            compressed_size,
        }
    }
}

/// Publishes `block` using `publisher` with a `subject_header` for header data and `subject_shard_prefix<SHARD_ID>`
/// for each shard data.
///
/// The shard subjects are padded with empty messages up to `shards`. The messages are compressed using zstd with
/// `compression_level` of -9 to 22.
pub async fn publish(
    block: impl Into<proto::Messages>,
    publisher: &impl Publisher,
    subject_header: impl Into<String>,
    subject_shard_prefix: impl Into<String>,
    compression_level: i32,
    shards: u64,
    last_msg_id: &mut Option<String>,
) -> Result<EncodingStats, PublishError> {
    let protobuf = block.into().into_inner();
    let height = protobuf
        .iter()
        .find_map(|msg| match msg.payload.as_ref().expect("Payload is mandatory") {
            proto::message::Payload::NearBlockHeader(header) => Some(header.header.as_ref().unwrap().height),
            _ => None,
        })
        .unwrap();
    let last_shard_id = protobuf.len() - 1;
    let subject_header = subject_header.into();
    let subject_shard_prefix = subject_shard_prefix.into();
    let shards = shards as usize;

    let mut encoded_size = 0;
    let mut compressed_size = 0;

    let messages = protobuf
        .into_iter()
        .filter_map(|msg| {
            let (msg_id, subject) = match msg.payload.as_ref().expect("Payload is mandatory") {
                proto::message::Payload::NearBlockHeader(..) => (height.to_string(), subject_header.clone()),
                proto::message::Payload::NearBlockShard(shard) => (
                    msg_id_for_shard(height, shard.shard_id),
                    subject_for_shard(&subject_shard_prefix, shard.shard_id),
                ),
                _ => return None,
            };

            let encoded = msg.encode_to_vec();
            encoded_size += encoded.len();

            let compressed = zstd::stream::encode_all(&encoded[..], compression_level).unwrap();
            compressed_size += compressed.len();

            Some((subject, msg_id, compressed.into()))
        })
        .chain((last_shard_id..shards).map(|shard_id| {
            let msg_id = msg_id_for_shard(height, shard_id as u64);
            let subject = subject_for_shard(&subject_shard_prefix, shard_id as u64);

            (subject, msg_id, Default::default())
        }));

    for (subject, msg_id, payload) in messages {
        let mut headers = HeaderMap::new();
        headers.insert(NATS_MESSAGE_ID, msg_id.as_str());
        if let Some(last_message_id) = last_msg_id.replace(msg_id) {
            headers.insert(NATS_EXPECTED_LAST_MESSAGE_ID, last_message_id.as_str());
        }

        publisher
            .publish_with_headers(subject, headers, payload)
            .await
            .map_err(PublishError::from)?;
    }

    Ok(EncodingStats::new(encoded_size, compressed_size))
}

fn msg_id_for_shard(height: u64, shard_id: ShardId) -> String {
    format!("{}.{}", height, shard_id)
}

fn subject_for_shard(subject_shard_prefix: &str, shard_id: ShardId) -> String {
    format!("{}{}", subject_shard_prefix, shard_id)
}

#[cfg(test)]
mod tests {
    use super::*;
    use async_trait::async_trait;
    use aurora_near_primitives::types::AccountId;
    use aurora_refiner_types::near_block::{BlockView, IndexerBlockHeaderView, NEARBlock, Shard};
    use std::ops::Deref;
    use std::str::FromStr;
    use std::sync::RwLock;
    use test_case::test_case;

    #[derive(Debug, PartialEq)]
    struct Publish {
        subject: String,
        headers: HeaderMap,
        is_empty: bool,
    }

    impl Publish {
        pub fn with_payload(subject: impl Into<String>, headers: impl Into<HeaderMap>) -> Self {
            Self {
                subject: subject.into(),
                headers: headers.into(),
                is_empty: false,
            }
        }

        pub fn without_payload(subject: impl Into<String>, headers: impl Into<HeaderMap>) -> Self {
            Self {
                subject: subject.into(),
                headers: headers.into(),
                is_empty: true,
            }
        }
    }

    #[derive(Debug)]
    struct DummyPublisher {
        messages: RwLock<Vec<Publish>>,
    }

    #[async_trait]
    impl Publisher for DummyPublisher {
        async fn publish_with_headers(
            &self,
            subject: String,
            headers: HeaderMap,
            payload: Bytes,
        ) -> Result<(), Box<dyn std::error::Error + Send + Sync>> {
            self.messages.write().unwrap().push(Publish {
                subject,
                headers,
                is_empty: payload.is_empty(),
            });
            Ok(())
        }
    }

    #[derive(Debug)]
    struct FailingPublisher(&'static str);

    #[async_trait]
    impl Publisher for FailingPublisher {
        async fn publish_with_headers(
            &self,
            _subject: String,
            _headers: HeaderMap,
            _payload: Bytes,
        ) -> Result<(), Box<dyn std::error::Error + Send + Sync>> {
            Ok(Err(self.0.to_owned())?)
        }
    }

    #[test_case(4, 4; "Publishing a block on all shards")]
    #[test_case(4, 1; "Publishing a block with some shards missing")]
    #[tokio::test]
    async fn test_publishing_block_succeeds(shards_count: u64, shards_filled: u64) {
        let block = create_dummy_block(shards_filled);
        let publisher = DummyPublisher {
            messages: RwLock::new(Vec::new()),
        };
        let subject_header = "header".to_owned();
        let subject_shard_prefix = "shard.".to_owned();
        let mut last_msg_id = None;

        publish(
            block,
            &publisher,
            subject_header,
            subject_shard_prefix,
            0,
            shards_count,
            &mut last_msg_id,
        )
        .await
        .unwrap();

        let messages = publisher.messages.read().unwrap();

        let actual_messages = messages.deref();
        let expected_messages = (0..shards_filled)
            .map(|shard_id| {
                Publish::with_payload(subject_for_shard("shard.", shard_id), {
                    let mut headers = HeaderMap::new();
                    headers.insert(NATS_MESSAGE_ID, msg_id_for_shard(0, shard_id).as_str());
                    if shard_id > 0 {
                        headers.insert(
                            NATS_EXPECTED_LAST_MESSAGE_ID,
                            msg_id_for_shard(0, shard_id - 1).as_str(),
                        );
                    }
                    headers
                })
            })
            .chain(std::iter::once(Publish::with_payload("header", {
                let mut headers = HeaderMap::new();
                headers.insert(NATS_MESSAGE_ID, "0");
                headers.insert(
                    NATS_EXPECTED_LAST_MESSAGE_ID,
                    msg_id_for_shard(0, shards_filled - 1).as_str(),
                );
                headers
            })))
            .chain((shards_filled..shards_count).map(|shard_id| {
                Publish::without_payload(subject_for_shard("shard.", shard_id), {
                    let mut headers = HeaderMap::new();
                    headers.insert(NATS_MESSAGE_ID, msg_id_for_shard(0, shard_id).as_str());
                    headers.insert(
                        NATS_EXPECTED_LAST_MESSAGE_ID,
                        if shard_id == shards_filled {
                            "0".to_owned()
                        } else {
                            msg_id_for_shard(0, shard_id - 1)
                        }
                        .as_str(),
                    );
                    headers
                })
            }))
            .collect::<Vec<_>>();

        assert_eq!(&expected_messages, actual_messages);
    }

    #[tokio::test]
    async fn test_publishing_block_on_failing_publisher_returns_error() {
        let shards = 1;
        let block = create_dummy_block(shards);
        let error_msg = "error";
        let publisher = FailingPublisher(error_msg);
        let subject_header = "header".to_owned();
        let subject_shard_prefix = "shard.".to_owned();
        let mut last_msg_id = None;

        let actual_error = publish(
            block,
            &publisher,
            subject_header,
            subject_shard_prefix,
            0,
            shards,
            &mut last_msg_id,
        )
        .await
        .expect_err("This must call FailingPublisher and fail");

        let expected_error = PublishError {
            message: error_msg.to_owned(),
        };

        assert_eq!(expected_error, actual_error);
    }

    fn create_dummy_block(shards_count: u64) -> NEARBlock {
        NEARBlock {
            block: create_dummy_block_header(shards_count),
            shards: (0..shards_count).map(create_dummy_shard).collect(),
        }
    }

    fn create_dummy_block_header(shards_count: u64) -> BlockView {
        BlockView {
            author: AccountId::from_str("dummy").unwrap(),
            header: IndexerBlockHeaderView {
                height: 0,
                prev_height: None,
                epoch_id: Default::default(),
                next_epoch_id: Default::default(),
                hash: Default::default(),
                prev_hash: Default::default(),
                prev_state_root: Default::default(),
                chunk_receipts_root: Default::default(),
                chunk_headers_root: Default::default(),
                chunk_tx_root: Default::default(),
                outcome_root: Default::default(),
                chunks_included: shards_count,
                challenges_root: Default::default(),
                timestamp: 0,
                timestamp_nanosec: 0,
                random_value: Default::default(),
                validator_proposals: vec![],
                chunk_mask: vec![],
                gas_price: 0,
                block_ordinal: None,
                total_supply: 0,
                challenges_result: vec![],
                last_final_block: Default::default(),
                last_ds_final_block: Default::default(),
                next_bp_hash: Default::default(),
                block_merkle_root: Default::default(),
                epoch_sync_data_hash: None,
                approvals: vec![],
                signature: Default::default(),
                latest_protocol_version: 0,
            },
        }
    }

    fn create_dummy_shard(shard_id: u64) -> Shard {
        Shard {
            shard_id,
            chunk: None,
            receipt_execution_outcomes: vec![],
            state_changes: vec![],
        }
    }
}
