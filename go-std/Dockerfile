FROM golang AS builder

WORKDIR /opt/src

COPY . /opt/src

RUN GOOS=linux CGO_ENABLED=0 go build -o main main.go

FROM scratch

EXPOSE 8080
ENV BENCH_HOST=0.0.0.0
ENV BENCH_PORT=8080
ENV BENCH_WORKERS=1
ENV GOMAXPROCS=$BENCH_WORKERS

WORKDIR /opt
COPY --from=builder /opt/src/main /opt/main

ENTRYPOINT ["./main"]
