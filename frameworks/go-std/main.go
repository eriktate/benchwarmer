package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
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
	Msg string `json:"msg"`
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, World!"))
}

func handleJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var req Greeting
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
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
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("failed to generate greeting"))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func handleDB(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var greeting Greeting
	if err := json.NewDecoder(r.Body).Decode(&greeting); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("failed to read request body"))
		return
	}
	defer r.Body.Close()

	id, err := CreateGreeting(db, greeting)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("failed to create greeting"))
		return
	}

	greeting, err = GetGreeting(db, id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("failed to get greeting"))
		return
	}

	data, err := json.Marshal(greeting)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("failed to marshal greeting"))
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
	case "/db":
		if r.Method != http.MethodPost {
			unsupported(w)
			return
		}
		handleDB(w, r)
	}
}

func run(addr string) error {
	return http.ListenAndServe(addr, http.HandlerFunc(handleRoutes))
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
