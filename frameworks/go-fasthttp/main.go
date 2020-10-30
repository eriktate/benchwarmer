package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/valyala/fasthttp"
)

type Greeting struct {
	ID        string    `json:"id,omitempty"`
	Greeting  string    `json:"greeting"`
	Name      string    `json:"name"`
	Used      bool      `json:"used"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type JSONRes struct {
	Greeting string `json:"msg"`
}

func handleHello(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	_, _ = ctx.Write([]byte("Hello, World!"))
}

func handleJSON(ctx *fasthttp.RequestCtx) {
	var req Greeting
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

func handleDB(ctx *fasthttp.RequestCtx) {
	var greeting Greeting
	if err := json.Unmarshal(ctx.PostBody(), &greeting); err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		_, _ = ctx.Write([]byte("failed to read request body"))
		return
	}

	id, err := CreateGreeting(db, greeting)
	if err != nil {
		log.Println(err)
		ctx.SetStatusCode(http.StatusInternalServerError)
		_, _ = ctx.Write([]byte("failed to create greeting"))
		return
	}

	greeting, err = GetGreeting(db, id)
	if err != nil {
		log.Println(err)
		ctx.SetStatusCode(http.StatusInternalServerError)
		_, _ = ctx.Write([]byte("failed to get greeting"))
		return
	}

	data, err := json.Marshal(greeting)
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
	case "/db":
		if !ctx.IsPost() {
			unsupported(ctx)
			return
		}
		handleDB(ctx)
	}
}

func run(addr string) error {
	return fasthttp.ListenAndServe(addr, handleRoutes)
}

func main() {
	var err error
	db, err = Connect()
	if err != nil {
		log.Fatal(err)
	}

	addr := fmt.Sprintf("%s:%s", os.Getenv("BENCH_HOST"), os.Getenv("BENCH_PORT"))
	log.Printf("GOMAXPROCS: %s", os.Getenv("GOMAXPROCS"))
	log.Printf("Listening on %s", addr)
	if err := run(addr); err != nil {
		log.Fatalf("server failed: %s", err)
	}
}
