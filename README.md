# borealis-prototypes
Prototype definitions for borealis messages

## Migration

### Example

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
