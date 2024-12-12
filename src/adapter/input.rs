use serde::Deserialize;

#[derive(Deserialize)]
pub struct AddNovelInput {
    name: String,
    finished: bool,
    chapter: u32,
}

pub fn adapt_add_novel_input(name: String, finished: bool, chapter: u32) -> AddNovelInput {
    AddNovelInput {
        name,
        finished,
        chapter,
    }
}
