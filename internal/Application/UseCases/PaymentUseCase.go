package usecases

import (
	entities "tech-challenge/internal/Domain/Entities"
	repositories "tech-challenge/internal/Domain/Repositories"
	"time"
)

type PaymentUseCase struct {
	repo *repositories.PaymentRepository
}

func NewPaymentUseCase(repo *repositories.PaymentRepository) *PaymentUseCase {
	return &PaymentUseCase{
		repo: repo,
	}
}

func (s *PaymentUseCase) GetPayments() ([]entities.Payment, error) {
	return s.repo.GetPayments()
}

func (s *PaymentUseCase) CreatePayment(payment entities.Payment) (entities.Payment, error) {
	payment.CreatedAt = time.Now()
	return s.repo.CreatePayment(payment)
}

func (s *PaymentUseCase) GetPayment(id string) (entities.Payment, error) {
	return s.repo.GetPayment(id)
}

func (s *PaymentUseCase) UpdatePayment(id string, updatedPayment entities.Payment) (entities.Payment, error) {
	return s.repo.UpdatePayment(id, updatedPayment)
}

func (s *PaymentUseCase) DeletePayment(id string) error {
	return s.repo.DeletePayment(id)
}
