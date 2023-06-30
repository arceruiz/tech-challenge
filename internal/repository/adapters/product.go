package adapters

import (
	"database/sql"
	"tech-challenge/internal/canonical"
	"tech-challenge/internal/repository"
	"tech-challenge/internal/repository/port"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepo() port.ProductRepository {
	return &productRepository{repository.New()}
}

func (r *productRepository) GetProducts() ([]canonical.Product, error) {
	return nil, nil
}

func (r *productRepository) CreateProduct(product canonical.Product) (canonical.Product, error) {
	return canonical.Product{}, nil
}

func (r *productRepository) UpdateProduct(id string, updatedProduct canonical.Product) (canonical.Product, error) {
	return canonical.Product{}, nil
}

func (r *productRepository) GetByID(id string) (canonical.Product, error) {
	return canonical.Product{}, nil
}

func (r *productRepository) GetByCategory(id string) ([]canonical.Product, error) {
	return nil, nil
}
