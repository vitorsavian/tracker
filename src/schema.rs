// @generated automatically by Diesel CLI.

diesel::table! {
    novels (id) {
        id -> Int4,
        title -> Varchar,
        chapter -> Int4,
    }
}
