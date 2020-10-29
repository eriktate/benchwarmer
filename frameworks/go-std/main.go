package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type JSONReq struct {
	Greeting string `json:"greeting"`
	Name     string `json:"name"`
}

type JSONRes struct {
	Msg string `json:"msg"`
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, World!"))
}

func handleJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var req JSONReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("failed to read request body"))
		return
	}
	defer r.Body.Close()

	res := JSONRes{
		Msg: fmt.Sprintf("%s %s", req.Greeting, req.Name),
	}

	data, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("failed to create greeting"))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func unsupported(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	_, _ = w.Write(nil)
}

func handleRoutes(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/hello":
		if r.Method != http.MethodGet {
			unsupported(w)
			return
		}
		handleHello(w, r)
	case "/json":
		if r.Method != http.MethodPost {
			unsupported(w)
			return
		}
		handleJSON(w, r)
	}
}

func run(addr string) error {
	return http.ListenAndServe(addr, http.HandlerFunc(handleRoutes))
}

func main() {
	addr := fmt.Sprintf("%s:%s", os.Getenv("BENCH_HOST"), os.Getenv("BENCH_PORT"))
	log.Printf("GOMAXPROCS: %s", os.Getenv("GOMAXPROCS"))
	log.Printf("Listening on %s", addr)
	if err := run(addr); err != nil {
		log.Fatalf("server failed: %s", err)
	}
}
