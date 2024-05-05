package main

import (
	"fmt"
	"net/http"
	"pong/handler"
)

func main() {
	handler := handler.NewPongHandler()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /pong", handler.HandlePong)

	fmt.Println("Starting pong server at :8002")
	http.ListenAndServe(":8002", mux)
}
