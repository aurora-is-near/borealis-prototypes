include!(concat!(env!("OUT_DIR"), "/_.rs"));

#[derive(Debug)]
pub struct Messages(Vec<Message>);

impl From<Vec<Message>> for Messages {
    fn from(value: Vec<Message>) -> Self {
        Messages(value)
    }
}

impl Messages {
    pub fn into_inner(self) -> Vec<Message> {
        self.0
    }
}
