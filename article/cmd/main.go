package main

import (
	"log"

	"github.com/yuhari7/article_service/api"
)

func main() {
	// Initialize the server
	e := api.NewServer()

	// Start the server
	log.Println("Server starting on :8001...")
	e.Logger.Fatal(e.Start(":8001"))
}
