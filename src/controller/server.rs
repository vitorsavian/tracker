use diesel::PgConnection;

pub trait App {
    fn database_connection(s: String) -> PgConnection;

    fn create_novel();
}
