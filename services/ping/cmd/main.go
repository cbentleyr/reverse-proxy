package main

import (
	"fmt"
	"net/http"
	"ping/handler"
)

func main() {
	handler := handler.NewPingHandler()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", handler.HandlePing)

	fmt.Println("Starting ping server at :8001")
	http.ListenAndServe(":8001", mux)
}
