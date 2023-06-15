package usecases

import (
	entities "tech-challenge/internal/Domain/Entities"
	repositories "tech-challenge/internal/Domain/Repositories"
)

type ProductUseCase struct {
	repo *repositories.ProductRepository
}

func NewProductUseCase(repo *repositories.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		repo: repo,
	}
}

func (s *ProductUseCase) GetProducts() ([]entities.Product, error) {
	return s.repo.GetProducts()
}

func (s *ProductUseCase) CreateProduct(product entities.Product) (entities.Product, error) {
	return s.repo.CreateProduct(product)
}

func (s *ProductUseCase) GetProduct(id string) (entities.Product, error) {
	return s.repo.GetProduct(id)
}

func (s *ProductUseCase) UpdateProduct(id string, updatedProduct entities.Product) (entities.Product, error) {
	return s.repo.UpdateProduct(id, updatedProduct)
}

func (s *ProductUseCase) DeleteProduct(id string) error {
	return s.repo.DeleteProduct(id)
}
