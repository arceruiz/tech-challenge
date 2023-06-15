package repositories

import (
	"database/sql"
	entities "tech-challenge/internal/Domain/Entities"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetProducts() ([]entities.Product, error) {
	return nil, nil
}

func (r *ProductRepository) CreateProduct(product entities.Product) (entities.Product, error) {
	return entities.Product{}, nil
}

func (r *ProductRepository) GetProduct(id string) (entities.Product, error) {
	return entities.Product{}, nil
}

func (r *ProductRepository) UpdateProduct(id string, updatedProduct entities.Product) (entities.Product, error) {
	return entities.Product{}, nil
}

func (r *ProductRepository) DeleteProduct(id string) error {
	return nil
}
