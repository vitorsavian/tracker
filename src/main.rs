pub mod routes;
use actix_web::{web, App, HttpServer};

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new().service(
            web::scope("/api").service(
                web::scope("/novel")
                    .service(routes::create_novel)
                    .service(routes::get_novels)
                    .service(routes::get_novel_by_id)
                    .service(routes::update_novel)
                    .service(routes::delete_novels),
            ),
        )
    })
    .bind(("127.0.0.1", 8080))?
    .run()
    .await
}
