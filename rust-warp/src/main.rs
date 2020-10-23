use std::env;
use std::net::{IpAddr, Ipv4Addr, SocketAddr};
use warp::Filter;

#[tokio::main(max_threads = 10_000)]
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
    warp::serve(hello).run(addr).await;
}
