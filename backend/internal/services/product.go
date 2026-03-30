package services

import (
	"github.com/JacobBananalDev/lacylorel/backend/internal/models"
	"github.com/JacobBananalDev/lacylorel/backend/internal/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) GetProducts() ([]models.Product, error) {
	return s.repo.GetProducts()
}
