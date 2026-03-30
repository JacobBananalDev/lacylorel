package repository

import (
	"time"

	"github.com/JacobBananalDev/lacylorel/backend/internal/models"
)

type ProductRepository interface {
	GetProducts() ([]models.Product, error)
}

type productRepository struct{}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (r *productRepository) GetProducts() ([]models.Product, error) {
	// TODO: replace with database query

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

	return products, nil
}