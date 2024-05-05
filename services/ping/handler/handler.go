package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PingHandler struct {
}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

type ResponseBody struct {
	Message string `json:"message"`
}

func (ph PingHandler) HandlePing(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Pinged!")
	body, _ := json.Marshal(ResponseBody{Message: "ping"})
	w.WriteHeader(200)
	w.Write(body)
}
