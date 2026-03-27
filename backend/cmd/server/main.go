package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type HealthCheckResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Set the content-type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Create the response date
	response := HealthCheckResponse{
		Status:  "UP",
		Message: "Service is healthy",
	}

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	// Register the health check handler for "/health" endpoint
	http.HandleFunc("/health", healthCheckHandler)

	// Start the server on port 8080
	log.Printf("Server starting on port 8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("HTTP server failed: ", err)
	}

}
