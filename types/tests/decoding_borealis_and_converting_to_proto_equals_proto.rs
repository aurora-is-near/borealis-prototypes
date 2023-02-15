use borealis_proto_types::{CompressedMessage, Message, Messages};
use borealis_types::message;
use borealis_types::payloads::NEARBlock;
use std::iter::once;
use test_case::test_case;

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
fn test_decoding_borealis_and_converting_to_proto_equals_proto(
    header_v3: &[u8],
    shards_v3: Vec<&[u8]>,
    block_v2: &[u8],
) {
    let expected_messages: Messages = shards_v3
        .into_iter()
        .map(|shard_v3| Message::decode_compressed(&shard_v3[..]).expect("Cannot decode compressed shard"))
        .chain(once(
            Message::decode_compressed(&header_v3[..]).expect("Cannot decode compressed header"),
        ))
        .collect();

    let actual_messages: Messages = message::Message::<NEARBlock>::from_cbor(block_v2)
        .expect("Cannot decode expected data from CBOR")
        .into();

    assert_eq!(expected_messages, actual_messages);
}