use borealis_proto_types::{CompressedMessage, Message, Messages};
use borealis_types::message;
use borealis_types::payloads::{NEARBlock, SendablePayload};
use std::iter::once;
use std::{fs, io};
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
fn test_decoding_proto_and_converting_to_borealis_equals_borealis(
    header_v3: &[u8],
    shards_v3: Vec<&[u8]>,
    block_v2: &[u8],
) {
    let mut expected_payload = message::Message::<NEARBlock>::from_cbor(block_v2)
        .expect("Cannot decode expected data from CBOR")
        .payload;

    upgrade_old_block(&mut expected_payload);
    dump_block(&expected_payload, "expected").unwrap();

    let messages: Messages = shards_v3
        .into_iter()
        .map(|shard_v3| Message::decode_compressed(&shard_v3[..]).expect("Cannot decode compressed shard"))
        .chain(once(
            Message::decode_compressed(&header_v3[..]).expect("Cannot decode compressed header"),
        ))
        .collect();

    let actual_payload = message::Message::<NEARBlock>::from(messages).payload;

    dump_block(&actual_payload, "actual").unwrap();

    let actual_payload = actual_payload.to_cbor().expect("Unable to encode NEAR block to CBOR");
    let expected_payload = expected_payload.to_cbor().expect("Unable to encode NEAR block to CBOR");

    assert_eq!(expected_payload, actual_payload);
}

/// Writes the contents of the `block` in a pretty-printed debug format into the test output tmp directory.
///
/// The directory is the one given from the env variable `CARGO_TARGET_TMPDIR`. The `block` height is used as a prefix
/// of the filename.
fn dump_block(block: &NEARBlock, suffix: &str) -> io::Result<()> {
    let contents = format!("{:#?}", block);
    let path = format!(
        "{}/{}.{}.txt",
        env!("CARGO_TARGET_TMPDIR"),
        block.block.header.height,
        suffix
    );
    fs::write(path, &contents.as_bytes())
}

/// Performs backwards compatible changes on the `block`.
///
/// The purpose of this change is to be able to compare the `block` historically encoded exactly with a newly encoded
/// one.
fn upgrade_old_block(block: &mut NEARBlock) {
    for shard in &mut block.shards {
        for receipt in &mut shard.receipt_execution_outcomes {
            if let Some(gas_profile) = &mut receipt.execution_outcome.outcome.metadata.gas_profile {
                for cost in gas_profile {
                    cost.cost = match cost.cost.as_str() {
                        "CONTRACT_COMPILE_BASE" => "CONTRACT_LOADING_BASE".to_owned(),
                        "CONTRACT_COMPILE_BYTES" => "CONTRACT_LOADING_BYTES".to_owned(),
                        "VALUE_RETURN" => "NEW_DATA_RECEIPT_BYTE".to_owned(),
                        _ => continue,
                    }
                }
            }
        }
    }
}
