package repository

import (
	"context"

	"github.com/JacobBananalDev/lacylorel/backend/internal/models"
	"github.com/jackc/pgx/v5"
)

type ProductRepository interface {
	GetProducts() ([]models.Product, error)
}

type productRepository struct {
	db *pgx.Conn
}

func NewProductRepository(db *pgx.Conn) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) GetProducts() ([]models.Product, error) {
	rows, err := r.db.Query(context.Background(), `
		SELECT id, name, description, price, image_url, created_at, updated_at
		FROM products
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var p models.Product
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.ImageURL,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}
