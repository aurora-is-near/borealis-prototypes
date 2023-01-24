mod decoding;
mod encoding;
mod proto;

pub use proto::*;

#[cfg(test)]
mod tests {
    use super::*;
    use borealis_types::compression::*;
    use borealis_types::message::*;
    use borealis_types::payloads::events::*;
    use borealis_types::payloads::request_response::*;
    use borealis_types::payloads::*;
    use borealis_types::prelude::*;
    use borealis_types::types::*;
    use near_primitives::types::*;

    #[test]
    fn test_roundtrip() {}
}
