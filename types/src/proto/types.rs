//! Borealis message protobuf definitions.
//!
//! The types in this module are automatically generated during build time from the `.proto` definitions.
include!(concat!(env!("OUT_DIR"), "/_.rs"));

#[derive(Debug, Clone, PartialEq)]
pub struct Messages(Vec<Message>);

impl From<Vec<Message>> for Messages {
    fn from(value: Vec<Message>) -> Self {
        Messages(value)
    }
}

impl FromIterator<Message> for Messages {
    fn from_iter<T: IntoIterator<Item = Message>>(iter: T) -> Self {
        From::from(iter.into_iter().collect::<Vec<Message>>())
    }
}

impl Messages {
    pub fn into_inner(self) -> Vec<Message> {
        self.0
    }
}
