use actix_web::{delete, get, post, put, web, App, HttpServer, Responder};

#[get("/api/novel/{id}")]
async fn get_novels(id: web::Path<String>) -> impl Responder {
    println!("{id}");
    format!("NATHAN VIADO")
}

#[post("/api/novel")]
async fn create_novel() -> impl Responder {
    format!("NATHAN VIADO")
}

#[put("/api/novel/{id}")]
async fn update_novel(id: web::Path<String>) -> impl Responder {
    println!("{id}");
    format!("NATHAN VIADO")
}

#[delete("/api/novel/{id}")]
async fn delete_novels(id: web::Path<String>) -> impl Responder {
    println!("{id}");
    format!("Deleted Novel")
}

#[actix_web::main] // or #[tokio::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| App::new().service(create_novel))
        .bind(("127.0.0.1", 8080))?
        .run()
        .await
}
