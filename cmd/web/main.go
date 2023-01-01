package main

import (
	"go-websocket-chat/internal/handlers"
	"log"
	"net/http"
)

func main() {
	mux := routes()

	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()

	log.Println("Starting webserver on port 8081")

	_ = http.ListenAndServe(":8081", mux)
}
