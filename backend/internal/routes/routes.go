package routes

import (
	"net/http"

	"github.com/JacobBananalDev/lacylorel/backend/internal/handlers"
)

func RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// API v1 routes
	mux.HandleFunc("/api/v1/health", handlers.HealthCheckHandler)

	return mux
}