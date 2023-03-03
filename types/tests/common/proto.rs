use async_nats::HeaderMap;
use async_trait::async_trait;
use borealis_proto_types::Publisher;
use bytes::Bytes;
use std::sync::RwLock;

#[derive(Clone, Debug, PartialEq)]
pub struct Publish {
    pub subject: String,
    pub headers: HeaderMap,
    pub payload: Bytes,
}

#[derive(Debug)]
pub struct DummyPublisher {
    pub messages: RwLock<Vec<Publish>>,
}

#[async_trait]
impl Publisher for DummyPublisher {
    async fn publish_with_headers(
        &self,
        subject: String,
        headers: HeaderMap,
        payload: Bytes,
    ) -> Result<(), Box<dyn std::error::Error + Send + Sync>> {
        self.messages.write().unwrap().push(Publish {
            subject,
            headers,
            payload,
        });
        Ok(())
    }
}
