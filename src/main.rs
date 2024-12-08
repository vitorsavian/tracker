use actix_web::{delete, get, post, put, web, App, HttpServer, Responder};

#[post("/")]
async fn create_novel() -> impl Responder {
    format!("Novel created")
}

#[get("/{id}")]
async fn get_novel_by_id(id: web::Path<String>) -> impl Responder {
    println!("{id}");
    format!("got novel")
}

#[get("/")]
async fn get_novels() -> impl Responder {
    format!("got novel")
}

#[put("/{id}")]
async fn update_novel(id: web::Path<String>) -> impl Responder {
    println!("{id}");
    format!("Novel updated")
}

#[delete("/{id}")]
async fn delete_novels(id: web::Path<String>) -> impl Responder {
    println!("{id}");
    format!("Deleted Novel")
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new().service(
            web::scope("/api").service(
                web::scope("/novel")
                    .service(create_novel)
                    .service(get_novels)
                    .service(get_novel_by_id)
                    .service(update_novel)
                    .service(delete_novels),
            ),
        )
    })
    .bind(("127.0.0.1", 8080))?
    .run()
    .await
}
