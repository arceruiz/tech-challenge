package repository

import (
	"client/internal/canonical"
	"database/sql"

	"github.com/sirupsen/logrus"
)

type ProductRepository interface {
	GetProducts() ([]canonical.Product, error)
	CreateProduct(product canonical.Product) (canonical.Product, error)
	UpdateProduct(id string, updatedProduct canonical.Product) (canonical.Product, error)
	GetByID(id string) (canonical.Product, error)
	GetByCategory(id string) ([]canonical.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepo() ProductRepository {
	connStr := "host=localhost port=5432 dbname=fiap_tech_challenge user=postgres password=1234 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logrus.Fatal(err)
	}
	return &productRepository{db}
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
