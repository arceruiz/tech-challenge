package service

import (
	"context"
	"tech-challenge/internal/canonical"
	"tech-challenge/internal/repository"
)

type ProductService interface {
	GetProducts(context.Context) ([]canonical.Product, error)
	CreateProduct(context.Context, canonical.Product) error
	UpdateProduct(context.Context, string, canonical.Product) error
	GetByID(context.Context, string) (*canonical.Product, error)
	GetByCategory(context.Context, string) ([]canonical.Product, error)
	Remove(context.Context, string) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService() ProductService {
	return &productService{
		repository.NewProductRepo(),
	}
}

func (s *productService) GetProducts(ctx context.Context) ([]canonical.Product, error) {
	return s.repo.GetProducts(ctx)
}

func (s *productService) CreateProduct(ctx context.Context, product canonical.Product) error {
	return s.repo.CreateProduct(ctx, product)
}

func (s *productService) UpdateProduct(ctx context.Context, id string, updatedProduct canonical.Product) error {
	return s.repo.UpdateProduct(ctx, id, updatedProduct)
}

func (s *productService) GetByID(ctx context.Context, id string) (*canonical.Product, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *productService) GetByCategory(ctx context.Context, id string) ([]canonical.Product, error) {
	return s.repo.GetByCategory(ctx, id)
}

func (s *productService) Remove(ctx context.Context, id string) error {
	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if product == nil {
		return canonical.ErrorNotFound
	}
	product.Status = 1
	err = s.repo.UpdateProduct(ctx, id, *product)
	if err != nil {
		return err
	}
	return nil
}
