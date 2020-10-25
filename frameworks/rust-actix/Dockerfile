FROM rust AS builder

WORKDIR /opt

# Build cargo cache before anything else
RUN apt-get update \
	&& apt-get install -y musl-tools \
	&& rustup target add x86_64-unknown-linux-musl

RUN USER=root cargo new benchwarmer

WORKDIR /opt/benchwarmer
COPY Cargo.toml Cargo.lock ./
RUN cargo build --release

# Build source code
COPY src ./src
RUN cargo install --target x86_64-unknown-linux-musl --path .

RUN cargo build --release


FROM scratch

ENV BENCH_HOST=0.0.0.0
ENV BENCH_PORT=8080
ENV BENCH_WORKERS=1

WORKDIR /opt
COPY --from=builder /usr/local/cargo/bin/rust-actix ./rust-actix
ENTRYPOINT ["./rust-actix"]
