use diesel::PgConnection;

pub trait App {
    fn database_connection(s: String) -> PgConnection;

    fn create_novel(&self);
    fn update_novel(&self);
    fn delete_novel(&self);
    fn get_novel(&self);
}
