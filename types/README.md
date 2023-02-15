# borealis-proto-types

Defines the Borealis message types in `protobuf` and compiles them to native Rust types.

## Installation

Run the following Cargo command in your project directory:

```shell
cargo add --git ssh://git@github.com/aurora-is-near/borealis-prototypes.git borealis-proto-types
```

Or add the following line to your Cargo.toml:

```toml
borealis-proto-types = { git = "ssh://git@github.com/aurora-is-near/borealis-prototypes.git", version = "0.2.0" }
```

## Usage

The Rust types in this crate are considered pure DTOs - they are used for reading data and contain no behavior. They may be used for message **transfer** and **storage**.

The usual usecase is for you to define your own types rich in behavior based on some business scenario. Then implement a conversion logic for these and the types provided by this crate.

```rust
// The main entry-point for decoding.
let message = borealis_proto_types::Message::decode_compressed(/* payload */)?;
```

```rust
// The main entry-point for encoding.
let payload = /* message */.encode_compressed(9)?;
```

Conversion logic for the commonly used types from `borealis-types` and `aurora-refiner-types` is already implemented so that the change easily integrates into projects already using those types.

You may parse blocks from a single stream if it carries shard 0 data like so.

```rust
use borealis_proto_types::message::Payload;
use borealis_proto_types::{CompressedMessage, Message as ProtoMsg};

if let Some(Payload::NearBlockShard(shard)) = ProtoMsg::decode_compressed(&msg.payload[..])?.payload {
    // Parse into NEAR block from shard data
    // (shards contain header data without `validator_proposals`, `challenges_result` and `approvals`)
    // Usage for `NEARBlock` from the crate `aurora_refiner_types`
    let block = aurora_refiner_types::near_block::NEARBlock::from(shard);
    // Or use the one from `borealis_types`
    let block = borealis_types::payloads::NEARBlock::from(shard);
};
```

Block header and shard data for every shard are published each with its own subject separately. This means you may have to consume 2..5 different message streams to reconstruct the full block.

## Example

```rust
use aurora_refiner_types::near_block::NEARBlock;
use borealis_proto_types::message::Payload;
use borealis_proto_types::{CompressedMessage, Message as ProtoMsg};

let payload = match borealis_proto_types::Message::decode_compressed(&msg.payload[..])?.payload {
    Some(Payload::NearBlockShard(shard)) => NEARBlock::from(shard),
    v => return Err(anyhow!("Unexpected payload {v:?}")),
};
println!("Received block hash {}", payload.block.header.hash);
```

For more possible scenarios of encoding and decoding see [test cases](/types/tests).
