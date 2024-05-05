package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PongHandler struct {
}

func NewPongHandler() *PongHandler {
	return &PongHandler{}
}

type ResponseBody struct {
	Message string `json:"message"`
}

func (ph PongHandler) HandlePong(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ponged!")
	body, _ := json.Marshal(ResponseBody{Message: "pong"})
	w.WriteHeader(200)
	w.Write(body)
}
