package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/JacobBananalDev/lacylorel/backend/internal/services"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	service := services.NewProductService()

	products, err := service.GetProducts()

	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
