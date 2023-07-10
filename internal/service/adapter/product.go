package adapter

import (
	"tech-challenge/internal/canonical"
	"tech-challenge/internal/repository/adapters"
	repos "tech-challenge/internal/repository/port"
	services "tech-challenge/internal/service/port"

	"github.com/google/uuid"
)

type productService struct {
	repo repos.ProductRepository
}

func NewProductService() services.ProductService {
	return &productService{
		adapters.NewProductRepo(),
	}
}

func (s *productService) GetProducts() ([]canonical.Product, error) {
	return s.repo.GetProducts()
}

func (s *productService) CreateProduct(product canonical.Product) error {
	product.ID = uuid.NewString()
	return s.repo.CreateProduct(product)
}

func (s *productService) UpdateProduct(id string, updatedProduct canonical.Product) error {
	return s.repo.UpdateProduct(id, updatedProduct)
}

func (s *productService) GetByID(id string) (*canonical.Product, error) {
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
	if product == nil {
		return canonical.ErrorNotFound
	}
	product.Status = "INACTIVE"
	err = s.repo.UpdateProduct(id, *product)
	if err != nil {
		return err
	}
	return nil
}
