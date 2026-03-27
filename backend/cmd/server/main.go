package main

import (
	"log"
	"net/http"

	"github.com/JacobBananalDev/lacylorel/backend/internal/handlers"
)


func main() {
	// Register the health check handler for "/health" endpoint
	http.HandleFunc("/health", handlers.HealthCheckHandler)

	// Start the server on port 8080
	log.Printf("Server starting on port 8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("HTTP server failed: ", err)
	}
}
