package main

import (
	"log"
	"net/http"
)

func main() {
	mux := routes()

	log.Println("Starting webserver on port 8081")

	_ = http.ListenAndServe(":8081", mux)
}
