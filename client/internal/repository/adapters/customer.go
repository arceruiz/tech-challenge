package adapters

import (
	"client/internal/canonical"
	"client/internal/repository"
	"client/internal/repository/ports"
	"database/sql"
)

type customerRepository struct {
	db *sql.DB
}

func NewCustomerRepo() ports.CustomerRepository {
	return &customerRepository{repository.New()}
}

func (r *customerRepository) Create(user canonical.Customer) error {
	sqlStatement := "INSERT INTO CUSTOMER (id, name, email, password, document, createdAt) VALUES ($1, $2, $3, $4, $5, $6)"

	_, err := r.db.Exec(sqlStatement, user.Id, user.Name, user.Email, user.Password, user.Document, user.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *customerRepository) GetByEmail(email string) (*canonical.Customer, error) {
	rows, err := r.db.Query(
		"SELECT * FROM Customer WHERE Email = $1",
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
			&user.CreatedAt,
			&user.Document,
			&user.Email,
			&user.Password,
			&user.Name,
		); err != nil {
			return nil, err
		}
		return &user, nil
	}

	return &user, repository.ErrorNotFound
}
