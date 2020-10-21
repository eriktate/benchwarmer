package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/valyala/fasthttp"
)

type JSONReq struct {
	Greeting string `json:"greeting"`
	Name     string `json:"name"`
}

type JSONRes struct {
	Greeting string `json:"msg"`
}

func handleHello(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	_, _ = ctx.Write([]byte("Hello, World!"))
}

func handleJSON(ctx *fasthttp.RequestCtx) {
	var req JSONReq
	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		_, _ = ctx.Write([]byte("failed to read request body"))
		return
	}

	res := JSONRes{
		Greeting: fmt.Sprintf("%s %s", req.Greeting, req.Name),
	}

	data, err := json.Marshal(res)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		_, _ = ctx.Write([]byte("failed to create greeting"))
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	_, _ = ctx.Write(data)
}

func unsupported(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
	_, _ = ctx.Write(nil)
}

func handleRoutes(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())
	switch path {
	case "/hello":
		if !ctx.IsGet() {
			unsupported(ctx)
			return
		}
		handleHello(ctx)
	case "/json":
		if !ctx.IsPost() {
			unsupported(ctx)
			return
		}
		handleJSON(ctx)
	}
}

func run(addr string) error {
	return fasthttp.ListenAndServe(addr, handleRoutes)
}

func main() {
	addr := fmt.Sprintf("%s:%s", os.Getenv("BENCH_HOST"), os.Getenv("BENCH_PORT"))
	log.Printf("Listening on %s", addr)
	if err := run(addr); err != nil {
		log.Fatalf("server failed: %s", err)
	}
}
