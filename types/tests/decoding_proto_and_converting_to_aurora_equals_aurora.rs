use crate::common::aurora::{dump_block, upgrade_old_block};
use aurora_refiner_types::near_block;
use borealis_proto_types::{CompressedMessage, Message, Messages};
use borealis_rs::bus_message::BusMessage;
use borealis_rs::bus_serde::ToCbor;
use std::iter::once;
use test_case::test_case;

pub mod common;

#[test_case(
    include_bytes!("data/34834053.v3"),
    vec![include_bytes!("data/34834053.0.v3")],
    include_bytes!("data/34834053.v2")
    ; "First block in history"
)]
#[test_case(
    include_bytes!("data/59869714.v3"),
    vec![
        include_bytes!("data/59869714.0.v3"),
        include_bytes!("data/59869714.1.v3"),
        include_bytes!("data/59869714.2.v3"),
        include_bytes!("data/59869714.3.v3"),
    ],
    include_bytes!("data/59869714.v2")
    ; "Block containing CONTRACT_COMPILE_BASE / CONTRACT_COMPILE_BYTES cost"
)]
#[test_case(
    include_bytes!("data/82041501.v3"),
    vec![
        include_bytes!("data/82041501.0.v3"),
        include_bytes!("data/82041501.1.v3"),
        include_bytes!("data/82041501.2.v3"),
        include_bytes!("data/82041501.3.v3"),
    ],
    include_bytes!("data/82041501.v2")
    ; "Block containing DEPLOY_CONTRACT action"
)]
#[test_case(
    include_bytes!("data/84797909.v3"),
    vec![
        include_bytes!("data/84797909.0.v3"),
        include_bytes!("data/84797909.1.v3"),
        include_bytes!("data/84797909.2.v3"),
        include_bytes!("data/84797909.3.v3"),
    ],
    include_bytes!("data/84797909.v2")
    ; "One of the recent blocks"
)]
#[test_case(
    include_bytes!("data/84800642.v3"),
    vec![
        include_bytes!("data/84800642.0.v3"),
        include_bytes!("data/84800642.1.v3"),
        include_bytes!("data/84800642.2.v3"),
        include_bytes!("data/84800642.3.v3"),
    ],
    include_bytes!("data/84800642.v2")
    ; "Block containing NEW_DATA_RECEIPT_BYTE cost"
)]
#[test_case(
    include_bytes!("data/88200321.v3"),
    vec![
        include_bytes!("data/88200321.0.v3"),
        include_bytes!("data/88200321.1.v3"),
        include_bytes!("data/88200321.2.v3"),
        include_bytes!("data/88200321.3.v3"),
    ],
    include_bytes!("data/88200321.v2")
    ; "Block containing NEW_ACTION_RECEIPT cost"
)]
fn test_decoding_proto_and_converting_to_aurora_equals_aurora(
    header_v3: &[u8],
    shards_v3: Vec<&[u8]>,
    block_v2: &[u8],
) {
    let mut expected_payload = BusMessage::<near_block::NEARBlock>::deserialize(block_v2)
        .expect("Cannot decode expected data from CBOR")
        .payload;

    upgrade_old_block(&mut expected_payload);
    dump_block(&expected_payload, "expected").unwrap();

    let messages: Messages = shards_v3
        .into_iter()
        .map(|shard_v3| Message::decode_compressed(shard_v3).expect("Cannot decode compressed shard"))
        .chain(once(
            Message::decode_compressed(header_v3).expect("Cannot decode compressed header"),
        ))
        .collect();

    let actual_payload = near_block::NEARBlock::from(messages);

    dump_block(&actual_payload, "actual").unwrap();

    let actual_payload = actual_payload.to_cbor().expect("Unable to encode NEAR block to CBOR");
    let expected_payload = expected_payload.to_cbor().expect("Unable to encode NEAR block to CBOR");

    assert_eq!(expected_payload, actual_payload);
}
