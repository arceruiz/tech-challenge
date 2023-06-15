package repositories

import (
	"database/sql"
	entities "tech-challenge/internal/Domain/Entities"
)

type PaymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

func (r *PaymentRepository) GetPayments() ([]entities.Payment, error) {
	return nil, nil
}

func (r *PaymentRepository) CreatePayment(payment entities.Payment) (entities.Payment, error) {
	return entities.Payment{}, nil
}

func (r *PaymentRepository) GetPayment(id string) (entities.Payment, error) {
	return entities.Payment{}, nil
}

func (r *PaymentRepository) UpdatePayment(id string, updatedPayment entities.Payment) (entities.Payment, error) {
	return entities.Payment{}, nil
}

func (r *PaymentRepository) DeletePayment(id string) error {
	return nil
}
