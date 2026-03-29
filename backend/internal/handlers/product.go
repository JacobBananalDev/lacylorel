package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JacobBananalDev/lacylorel/backend/internal/models"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	products := []models.Product{
		{
			ID:          1,
			Name:        "Product 1",
			Description: "Description for product 1",
			Price:       1000,
			ImageURL:    "https://example.com/product1.jpg",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Name:        "Product 2",
			Description: "Description for product 2",
			Price:       2000,
			ImageURL:    "https://example.com/product2.jpg",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
