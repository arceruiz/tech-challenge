package service

import (
	"context"
	"errors"
	"tech-challenge/internal/canonical"
	"tech-challenge/internal/repository"
)

type PaymentService interface {
	GetByID(context.Context, string) (*canonical.Payment, error)
	Callback(ctx context.Context, paymentId, status string) error
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

func (s *paymentService) Callback(ctx context.Context, paymentId, status string) error {
	payment, err := s.repo.GetByID(ctx, paymentId)
	if err != nil {
		return err
	}
	if payment == nil {
		return errors.New("payment not found")
	}
	payment.Status = s.statusToCannonical(status)
	err = s.repo.Update(ctx, paymentId, *payment)
	if err != nil {
		return err
	}
	return nil
}

func (s *paymentService) statusToCannonical(status string) canonical.PaymentStatus {
	if val, ok := canonical.MapPaymentStatus[status]; ok {
		return val
	}
	return canonical.PAYMENT_ERROR
}
