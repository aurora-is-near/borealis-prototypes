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
    fs::write(path, contents.as_bytes())
}
