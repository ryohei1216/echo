package main

import (
	"encoding/json"
	"log"
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

	jsonData, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
