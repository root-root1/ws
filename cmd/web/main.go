package main

import (
	"github.com/root-root1/ws/internal/handlers"
	"log"
	"net/http"
)

func main() {
	mux := routes()

	log.Println("Starting Channels Listener")
	go handlers.ListenToWsChannel()

	log.Println("Server is up and Running on http://localhost:8000")
	_ = http.ListenAndServe(":8000", mux)
}
