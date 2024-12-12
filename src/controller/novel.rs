use core::panic;

use diesel::{Connection, PgConnection};

use crate::controller::server::App;

pub struct Controller {
    pub conn: PgConnection,
}

impl App for Controller {
    fn database_connection(s: String) -> PgConnection {
        let conn =
            PgConnection::establish(&s).unwrap_or_else(|_| panic!("Error connecting to {}", s));

        conn
    }

    fn create_novel() {}
    // add code here
}
