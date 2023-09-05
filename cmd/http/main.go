package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/ping", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, _ *http.Request) {
	res := Response{Message: "hello"}

	jsonData, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// レスポンスヘッダーの設定
	w.Header().Set("Content-Type", "application/json")

	// ステータスコードを設定
	w.WriteHeader(http.StatusOK)

	w.Write(jsonData)
}
