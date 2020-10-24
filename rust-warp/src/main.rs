use serde::{Deserialize, Serialize};
use std::env;
use std::net::{IpAddr, Ipv4Addr, SocketAddr};
use warp::Filter;

#[derive(Deserialize)]
struct JSONReq {
    greeting: String,
    name: String,
}

#[derive(Serialize)]
struct JSONRes {
    msg: String,
}

#[tokio::main]
async fn main() {
    // get env configs
    let host_env = env::var("BENCH_HOST").unwrap_or(String::from("127.0.0.1"));
    let port_env = env::var("BENCH_PORT").unwrap_or(String::from("8080"));

    println!("Listening on {}:{}", host_env, port_env);
    // make listen addr
    let host: Ipv4Addr = host_env.parse().unwrap();
    let port = port_env.parse::<u16>().unwrap();
    let addr = SocketAddr::new(IpAddr::V4(host), port);

    let hello = warp::path!("hello").map(|| format!("Hello, world!"));
    let json = warp::post()
        .and(warp::path("json"))
        .and(warp::body::json())
        .map(|req: JSONReq| {
            let res = JSONRes {
                msg: format!("{} {}", req.greeting, req.name),
            };
            warp::reply::json(&res)
        });

    let routes = hello.or(json);
    warp::serve(routes).run(addr).await;
}
