package service

import (
	"context"
	"tech-challenge/internal/canonical"
	"tech-challenge/internal/repository"
)

type PaymentService interface {
	GetByID(context.Context, string) (*canonical.Payment, error)
}

type paymentService struct {
	repo repository.PaymentRepository
}

func NewPaymentService() PaymentService {
	return &paymentService{
		repository.NewPaymentRepo(),
	}
}

func (s *paymentService) GetByID(ctx context.Context, id string) (*canonical.Payment, error) {
	return s.repo.GetByID(ctx, id)
}
