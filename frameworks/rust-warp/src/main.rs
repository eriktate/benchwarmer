use serde::{Deserialize, Serialize};
use std::env;
use std::net::{IpAddr, Ipv4Addr, SocketAddr};
use tokio::runtime;
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

async fn run() {
    // get env configs
    let host_env = env::var("BENCH_HOST").unwrap_or(String::from("127.0.0.1"));
    let port_env = env::var("BENCH_PORT").unwrap_or(String::from("8080"));

    println!("Listening on {}:{}", host_env, port_env);
    // make listen addr
    let host: Ipv4Addr = host_env.parse().unwrap();
    let port = port_env.parse::<u16>().unwrap();
    let addr = SocketAddr::new(IpAddr::V4(host), port);

    let hello = warp::path!("hello").map(|| format!("Hello, World!"));
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

fn main() {
    let workers_env = std::env::var("BENCH_WORKERS").unwrap_or(String::from("1"));
    let workers: usize = workers_env.parse().unwrap();
    // need to switch schedulers depending on the number of workers defined
    let mut runtime = if workers > 1 {
        runtime::Builder::new()
            .threaded_scheduler()
            .core_threads(workers)
            .enable_all()
            .build()
            .unwrap()
    } else {
        runtime::Builder::new()
            .basic_scheduler()
            .enable_all()
            .build()
            .unwrap()
    };

    runtime.block_on(async { run().await })
}
