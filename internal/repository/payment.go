package repository

import (
	"context"
	"tech-challenge/internal/canonical"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PaymentRepository interface {
	GetByID(context.Context, string) (*canonical.Payment, error)
}

type paymentRepository struct {
	db *pgxpool.Pool
}

func NewPaymentRepo() PaymentRepository {
	return &paymentRepository{New()}
}

func (r *paymentRepository) GetByID(ctx context.Context, id string) (*canonical.Payment, error) {
	rows, err := r.db.Query(ctx,
		"SELECT * FROM \"Payment\" WHERE ID = $1",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payment canonical.Payment
	if rows.Next() {
		if err = rows.Scan(
			&payment.ID,
			&payment.PaymentType,
			&payment.CreatedAt,
			&payment.Status,
		); err != nil {
			return nil, err
		}
	}

	return &payment, nil
}
