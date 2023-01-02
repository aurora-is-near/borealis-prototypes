use std::io::Result;

fn main() -> Result<()> {
    prost_build::compile_protos(
        &["proto/message.proto"],
        &[
            "proto/",
            "proto/payloads",
            "proto/payloads/near_block",
            "proto/payloads/near_block/transaction",
        ],
    )?;
    Ok(())
}
