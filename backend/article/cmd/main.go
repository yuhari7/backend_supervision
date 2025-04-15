package main

import (
	"log"

	"github.com/yuhari7/backend_supervision/article/api"
)

func main() {
	// Initialize and start the server
	server := api.NewServer()

	// Log that the server has started successfully
	log.Println("✅ Server started successfully at", server.Listener.Addr())
}
