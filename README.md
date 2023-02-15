# borealis-prototypes
Protobuf definitions for borealis messages, referred to as "version 3" borealis messages and are encoded and compressed differently than "version 2".

## Motivation

The borealis message types are becoming too large in size, so much so that they fail to be sent to jetstream in extreme cases. These extreme cases usually contained a lot of contract deployments in a single block. Consequentially, these might crash parts of the borealis infrastructure when encountered.

These types bring the size down by introducing Protobuf encoding, which is the best in terms of space & performance. Then a ZSTD compression is applied to bring the size down even futher.

The results are 33% average space reduction after switching to protobuf. The compression brings the size down by an additional third of the original size. Resulting in the message stream of one third the size of the original.

## Usage

This repository contains common definition of protobuf types and their support for multiple languages.

Furthermore, it contains a CLI tool for migrating the version 2 Borealis messages to version 3.

### Types

The common `.proto` definitions are contained in this directory.

```
types/proto
```

Pick a guide based on the language you are using to decode and encode the messages.

* [go](/go/README.md)
* [rust](/types/README.md)

### Migration

You need [rust](https://rustup.rs/) to use this.

Run the following command to see documentation for the migration command.

```shell
cargo run -p borealis-proto-migrate migrate help
````

The tool creates multiple output messages (2..5) per single input message. It does so in a way that they match in sequence number. Meaning that for messages with less than 4 shard blocks, there are empty messages published to keep the sequence number in sync.

## Example

### Run migration tool

Downloads version 2 messages from `--nats-server` and uploads them re-encoded and compressed as version 3 messages to `--nats-output-server`.

```shell
RUST_LOG=debug cargo run -p borealis-proto-migrate -- \
  --nats-creds production_developer.creds \
  --nats-server nats://developer.nats.backend.aurora.dev:4222 \
  --nats-subject v2.mainnet.near.blocks \
  --nats-output-server nats://0.0.0.0:4222 \
  --nats-output-subject-header v2.mainnet.near.blocks \
  --nats-output-subject-shards v2.mainnet.near. \
  migrate
```

## Data availability / developer streams

  - Server:  nats://developer.nats.backend.aurora.dev:4222/
  - Stream name: v3_mainnet_near_blocks
  - Subject header: v3.mainnet.near.blocks.header
  - Subject shards: v3.mainnet.near.blocks.[0-4]
