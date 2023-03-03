use crate::common::aurora::{dump_block, upgrade_old_block};
use crate::common::proto::DummyPublisher;
use aurora_refiner_types::near_block::NEARBlock;
use borealis_proto_types::{publish, BlocksBuilder};
use borealis_rs::bus_message::BusMessage;
use borealis_rs::bus_serde::ToCbor;
use std::ops::Deref;
use std::sync::RwLock;
use test_case::test_case;

pub mod common;

#[test_case(
    include_bytes!("data/34834053.v2")
    ; "First block in history"
)]
#[test_case(
    include_bytes!("data/59869714.v2")
    ; "Block containing CONTRACT_COMPILE_BASE / CONTRACT_COMPILE_BYTES cost"
)]
#[test_case(
    include_bytes!("data/82041501.v2")
    ; "Block containing DEPLOY_CONTRACT action"
)]
#[test_case(
    include_bytes!("data/84797909.v2")
    ; "One of the recent blocks"
)]
#[test_case(
    include_bytes!("data/84800642.v2")
    ; "Block containing NEW_DATA_RECEIPT_BYTE cost"
)]
#[tokio::test]
async fn test_published_messages_are_decoded_by_builder(block_v2: &[u8]) {
    let height = 0;
    let shards = 4;
    let mut expected_payload = BusMessage::<NEARBlock>::deserialize(block_v2)
        .expect("Cannot decode expected data from CBOR")
        .payload;

    upgrade_old_block(&mut expected_payload);
    dump_block(&expected_payload, "published.expected").unwrap();

    let mut last_msg_id: Option<String> = None;
    let publisher = DummyPublisher {
        messages: RwLock::new(Vec::new()),
    };

    publish(
        expected_payload.clone(),
        &publisher,
        "header",
        "shard.",
        0,
        shards,
        &mut last_msg_id,
    )
    .await
    .expect("Unable to publish");

    let mut builder = BlocksBuilder::default();

    let messages = publisher.messages.read().unwrap();
    let messages = messages.deref();
    let messages = messages.iter().cloned().map(|message| async_nats::Message {
        subject: message.subject,
        reply: None,
        payload: message.payload,
        headers: Some(message.headers),
        status: None,
        description: None,
        length: 0,
    });

    let mut actual_payload: Option<NEARBlock> = None;

    for message in messages {
        if let Some(payload) = builder.add_message(height, message).expect("Unable to add message") {
            actual_payload.replace(payload);
        }
    }

    let actual_payload = actual_payload.unwrap();

    dump_block(&actual_payload, "published.actual").unwrap();

    let actual_payload = actual_payload.to_cbor().expect("Unable to encode NEAR block to CBOR");
    let expected_payload = expected_payload.to_cbor().expect("Unable to encode NEAR block to CBOR");

    assert_eq!(expected_payload, actual_payload);
}
