package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	r := chi.NewRouter()
	r.Get("/ping", handler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handler(w http.ResponseWriter, _ *http.Request) {
	res := Response{Message: "hello"}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		slog.Warn(fmt.Sprintf("encode response: %v", err))
	}
}
