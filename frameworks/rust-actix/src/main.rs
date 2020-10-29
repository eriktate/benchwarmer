use actix_web::{get, post, web, App, HttpResponse, HttpServer, Responder};
use serde::{Deserialize, Serialize};
use std::env;

#[derive(Deserialize, Serialize)]
struct JSONReq {
    greeting: String,
    name: String,
}

#[derive(Deserialize, Serialize)]
struct JSONRes {
    msg: String,
}

#[post("/json")]
async fn json(req: web::Json<JSONReq>) -> impl Responder {
    let res = JSONRes {
        msg: format!("{} {}", req.greeting, req.name),
    };

    HttpResponse::Ok().json(&res)
}

#[get("/hello")]
async fn hello() -> impl Responder {
    String::from("Hello, World!")
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let host = env::var("BENCH_HOST").unwrap_or(String::from("127.0.0.1"));
    let port = env::var("BENCH_PORT").unwrap_or(String::from("8080"));
    let workers_env = env::var("BENCH_WORKERS").unwrap_or(String::from("4"));

    let workers = workers_env.parse::<usize>().unwrap();

    println!("Listening on {}:{}", host, port);
    HttpServer::new(|| App::new().service(hello).service(json))
        .bind(format!("{}:{}", host, port))?
        .workers(workers)
        .run()
        .await
}
