package main

import (
	"log"
	"net/http"

	"github.com/JacobBananalDev/lacylorel/backend/internal/routes"
)

func main() {
	// Register the routes
	router := routes.RegisterRoutes()

	log.Println("Server starting on port 8080...")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("HTTP server failed: ", err)
	}
}
