use aurora_refiner_types::near_block::NEARBlock;
use std::{fs, io};

/// Writes the contents of the `block` in a pretty-printed debug format into the test output tmp directory.
///
/// The directory is the one given from the env variable `CARGO_TARGET_TMPDIR`. The `block` height is used as a prefix
/// of the filename.
#[cfg(test)]
pub fn dump_block(block: &NEARBlock, suffix: &str) -> io::Result<()> {
    let contents = format!("{:#?}", block);
    let path = format!(
        "{}/{}.{suffix}.txt",
        env!("CARGO_TARGET_TMPDIR"),
        block.block.header.height,
    );
    fs::write(path, &contents.as_bytes())
}

/// Performs backwards compatible changes on the `block`.
///
/// The purpose of this change is to be able to compare the `block` historically encoded exactly with a newly encoded
/// one.
#[cfg(test)]
pub fn upgrade_old_block(block: &mut NEARBlock) {
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
