use actix_web::{delete, get, post, put, web, Responder};

struct Controller {}

impl Controller {
    // add code here
}

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
