package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/apex/gateway"
)

func main() {
	http.HandleFunc("/hello", hello)
	log.Fatal(gateway.ListenAndServe(":3000", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Go Serverless v1.0! Your function executed successfully!",
	})
}
