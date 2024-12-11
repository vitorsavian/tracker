use crate::controller::novel;
use actix_web::{web, App, HttpServer};

#[actix_web::main]
async fn run() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new().service(
            web::scope("/api").service(
                web::scope("/novel")
                    .service(novel::create_novel)
                    .service(novel::get_novels)
                    .service(novel::get_novel_by_id)
                    .service(novel::update_novel)
                    .service(novel::delete_novels),
            ),
        )
    })
    .bind(("127.0.0.1", 8080))?
    .run()
    .await
}
