//! This module is responsible for the construction of blocks from multiple [`async_nats`] messages.
//!
//! As pairing them and checking whether they're completely received yet is an expected problem when reading v3 streams,
//! it is considered generic and belongs to the library.
//!
//! The state is kept in [`BlocksBuilder`] object, you may construct it by calling [`BlocksBuilder::default`].
//!
//! See the documentation of [`BlocksBuilder`] to learn more.
//!
//! # Examples
//!
//! ```
//! use borealis_proto_types::BlocksBuilder;
//!
//! # fn main() {
//! let blocks = BlocksBuilder::default();
//! # }
//! ```
use crate::message::Payload::{NearBlockHeader, NearBlockShard};
use crate::{CompressedMessage, DecodeError, Message as ProtoMsg};
use async_nats::Message;
use aurora_refiner_types::near_block::{BlockView, NEARBlock, Shard};
use itertools::Itertools;
use std::collections::HashMap;

/// Collects messages, each one containing a part of a block data. There are multiple parts, one for a header and one
/// for each shard. All messages of a block need to be collected to construct a single [`NEARBlock`].
///
/// The blocks are identified by message sequence number. If there are blocks with less than four shards, an empty
/// message is expected at its place. Call [`skip_shard`] on that one.
///
/// If the message contains a payload, call [`add_message`] on it. This call may result in completing a [`NEARBlock`],
/// in case that happens it wraps it in [`Some`] and transfers ownership to the caller.
///
/// # Examples
///
/// ```
/// use async_nats::{header::NATS_MESSAGE_ID, Message};
/// use borealis_proto_types::{BlocksBuilder, message::Payload::{NearBlockHeader, NearBlockShard}};
///
/// # fn next_message() -> Option<Message> { None }
/// # fn main() -> Result<(), Box<dyn std::error::Error>> {
/// let mut blocks = BlocksBuilder::default();
///
/// while let Some(msg) = next_message() {
///     // Get height from msg, in this case from the NATS message header of format: `<height>[.<shard_id>]`
///     let height = msg.headers
///         .as_ref()
///         .expect("Messages must contain headers")
///         .get(NATS_MESSAGE_ID)
///         .expect("Message must contain NATS_MESSAGE_ID")
///         .to_string()
///         .split('.')
///         .next()
///         .unwrap()
///         .parse::<u64>()?;
///
///     // Add msg and get block if it is completed
///     match blocks.add_message(height, msg)? {
///         Some(near_block) => { /* Do something with completed near_block */ }
///         None => continue,
///     }
/// }
/// # Ok(())
/// # }
/// ```
///
/// [`skip_shard`]: BlocksBuilder::skip_shard
#[derive(Debug, Default)]
pub struct BlocksBuilder {
    /// Maps block height to the builder constructing the block of that height from the messages.
    blocks: HashMap<u64, BlockBuilder>,
}

impl BlocksBuilder {
    /// Adds the [`Message`] into the block builder identified by the `height`.
    ///
    /// Returns [`Some`] with the constructed block if all the messages (one for header and one for each shard) are
    /// collected. Returns [`None`] if the block is still incomplete.
    pub fn add_message(&mut self, height: u64, msg: Message) -> Result<Option<NEARBlock>, DecodeError> {
        Ok(if msg.payload.is_empty() {
            None
        } else {
            let decoded_message = ProtoMsg::decode_compressed(&msg.payload[..])?;

            self.add_proto_message(height, decoded_message)
        })
    }

    fn add_proto_message(&mut self, height: u64, msg: ProtoMsg) -> Option<NEARBlock> {
        let maybe_block = match msg.payload {
            Some(NearBlockHeader(header)) => self.blocks.entry(height).or_default().add_header(header.into()),
            Some(NearBlockShard(shard)) => self.blocks.entry(height).or_default().add_shard(shard.into()),
            _ => return None,
        }
        .build();

        if maybe_block.is_some() {
            self.blocks.remove(&height);
        }

        maybe_block
    }
}

#[derive(Debug, Default)]
struct BlockBuilder {
    header: Option<BlockView>,
    shards: HashMap<u64, Shard>,
}

impl BlockBuilder {
    fn add_header(&mut self, header: BlockView) -> &mut Self {
        self.header.replace(header);
        self
    }

    fn add_shard(&mut self, shard: Shard) -> &mut Self {
        self.shards.insert(shard.shard_id, shard);
        self
    }

    fn is_ready(&self) -> bool {
        self.header
            .as_ref()
            .map(|header| self.shards.len() as u64 >= header.header.chunks_included)
            .unwrap_or(false)
    }

    fn build(&mut self) -> Option<NEARBlock> {
        self.is_ready().then(|| NEARBlock {
            block: self.header.take().unwrap(),
            shards: self
                .shards
                .drain()
                .map(|(_, shard)| shard)
                .sorted_by_key(|shard| shard.shard_id)
                .collect(),
        })
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::proto;
    use aurora_refiner_types::near_block::IndexerBlockHeaderView;
    use near_primitives::types::AccountId;
    use std::str::FromStr;
    use test_case::test_case;

    #[test_case(0; "No shards")]
    #[test_case(1; "One shard")]
    #[test_case(4; "Four shards")]
    fn test_blocks_builder_builds_only_after_adding_header_and_expected_number_of_shards(shards_count: u64) {
        let height = 0;
        let mut builder = BlocksBuilder::default();

        let header = create_dummy_block_header_message(shards_count);

        let mut block = builder.add_message(height, header).unwrap();

        for i in 0..shards_count {
            assert!(block.is_none());

            let shard = create_dummy_shard_message(i);

            block = builder.add_message(height, shard).unwrap();
        }

        assert!(block.is_some());
    }

    #[test_case(0; "No shards")]
    #[test_case(1; "One shard")]
    #[test_case(4; "Four shards")]
    fn test_block_builder_builds_only_after_adding_header_and_expected_number_of_shards(shards_count: u64) {
        let mut builder = BlockBuilder::default();

        let header = create_dummy_block_header(shards_count);

        builder.add_header(header);

        for i in 0..shards_count {
            assert!(builder.build().is_none());

            let shard = create_dummy_shard(i);

            builder.add_shard(shard);
        }

        assert!(builder.build().is_some());
    }

    #[test]
    fn test_empty_block_builder_builds_no_block() {
        let mut builder = BlockBuilder::default();
        assert!(builder.build().is_none());
    }

    #[test]
    fn test_blocks_builder_can_repeatedly_build_blocks() {
        let shards_count = 0;
        let mut builder = BlocksBuilder::default();

        for height in 0..2 {
            let header = create_dummy_block_header_message(shards_count);
            let block = builder.add_message(height, header).unwrap();

            assert!(block.is_some());
        }
    }

    #[test]
    fn test_blocks_builder_can_build_blocks_out_of_order() {
        let shards_count = 1;
        let mut builder = BlocksBuilder::default();

        let height = 0;
        let header = create_dummy_block_header_message(shards_count);
        let block = builder.add_message(height, header).unwrap();

        assert!(block.is_none());

        let height = 1;
        let header = create_dummy_block_header_message(shards_count);
        let block = builder.add_message(height, header).unwrap();

        assert!(block.is_none());

        let height = 0;
        let header = create_dummy_shard_message(0);
        let block = builder.add_message(height, header).unwrap();

        assert!(block.is_some());

        let height = 1;
        let header = create_dummy_shard_message(0);
        let block = builder.add_message(height, header).unwrap();

        assert!(block.is_some());
    }

    #[test]
    fn test_adding_messages_with_empty_payload_to_blocks_builder_has_no_effect() {
        let empty_count = 3;
        let shards_count = 1;
        let height = 0;
        let mut builder = BlocksBuilder::default();

        let header = create_dummy_block_header_message(shards_count);

        let mut block = builder.add_message(height, header).unwrap();

        assert!(block.is_none());

        for _ in 0..empty_count {
            let empty = create_empty_message();

            block = builder.add_message(height, empty).unwrap();

            assert!(block.is_none());
        }

        let shard = create_dummy_shard_message(0);

        block = builder.add_message(height, shard).unwrap();

        assert!(block.is_some());
    }

    impl From<proto::Message> for Message {
        fn from(value: proto::Message) -> Self {
            let compression_level = 0;
            let payload = value.encode_compressed(compression_level).unwrap();

            Message {
                payload: payload.into(),
                ..create_empty_message()
            }
        }
    }

    fn create_empty_message() -> Message {
        Message {
            subject: "".to_string(),
            reply: None,
            payload: Default::default(),
            headers: None,
            status: None,
            description: None,
            length: 0,
        }
    }

    fn create_dummy_block_header_message(shards_count: u64) -> Message {
        proto::Message {
            payload: Some(NearBlockHeader(create_dummy_block_header(shards_count).into())),
        }
        .into()
    }

    fn create_dummy_shard_message(shard_id: u64) -> Message {
        proto::Message {
            payload: Some(NearBlockShard(create_dummy_shard(shard_id).into())),
        }
        .into()
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
