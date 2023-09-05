package repository

import (
	"context"
	"tech-challenge/internal/canonical"

	"github.com/jackc/pgx/v4/pgxpool"
)

type CustomerRepository interface {
	Create(context.Context, canonical.Customer) (int, error)
	GetByEmail(context.Context, string) (*canonical.Customer, error)
}

type customerRepository struct {
	db *pgxpool.Pool
}

func NewCustomerRepo() CustomerRepository {
	return &customerRepository{New()}
}

func (r *customerRepository) Create(ctx context.Context, user canonical.Customer) (int, error) {
	sqlStatement := "INSERT INTO \"Customer\" (name, email, password, document, createdAt) VALUES ($1, $2, $3, $4, $5) RETURNING ID"
	var insertedId int

	err := r.db.QueryRow(ctx, sqlStatement, user.Name, user.Email, user.Password, user.Document, user.CreatedAt).Scan(&insertedId)
	if err != nil {
		return 0, err
	}

	return insertedId, nil
}

func (r *customerRepository) GetByEmail(ctx context.Context, email string) (*canonical.Customer, error) {
	rows, err := r.db.Query(ctx,
		"SELECT * FROM \"Customer\" WHERE Email = $1",
		email,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var user canonical.Customer

	if rows.Next() {
		if err = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.Document,
			&user.Password,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		return &user, nil
	}

	return nil, ErrorNotFound
}
