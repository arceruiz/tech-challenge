package service

import (
	"client/internal/canonical"
	"client/internal/repository"
)

type ProductService interface {
	GetProducts() ([]canonical.Product, error)
	CreateProduct(product canonical.Product) (canonical.Product, error)
	UpdateProduct(id string, updatedProduct canonical.Product) (canonical.Product, error)
	GetByID(id string) (canonical.Product, error)
	GetByCategory(id string) ([]canonical.Product, error)
	Remove(id string) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService() ProductService {
	return &productService{
		repository.NewProductRepo(),
	}
}

func (s *productService) GetProducts() ([]canonical.Product, error) {
	return s.repo.GetProducts()
}

func (s *productService) CreateProduct(product canonical.Product) (canonical.Product, error) {
	return s.repo.CreateProduct(product)
}

func (s *productService) UpdateProduct(id string, updatedProduct canonical.Product) (canonical.Product, error) {
	return s.repo.UpdateProduct(id, updatedProduct)
}

func (s *productService) GetByID(id string) (canonical.Product, error) {
	return s.repo.GetByID(id)
}

func (s *productService) GetByCategory(id string) ([]canonical.Product, error) {
	return s.repo.GetByCategory(id)
}

func (s *productService) Remove(id string) error {
	product, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	product.Status = "REMOVED"
	_, err = s.repo.UpdateProduct(id, product)
	if err != nil {
		return err
	}
	return nil
}
