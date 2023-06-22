package repository

import (
	"client/internal/canonical"
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var (
	errorNotFound = errors.New("entity not found")
)

type Repository interface {
	Create(canonical.User) error
	GetByEmail(email string) (*canonical.User, error)
}

type repo struct {
	db *sql.DB
}

func New() Repository {
	connStr := "host=localhost port=5432 dbname=fiap_tech_challenge user=postgres password=1234 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logrus.Fatal(err)
	}
	return &repo{db}
}

func (r *repo) Create(user canonical.User) error {
	sqlStatement := "INSERT INTO CUSTOMER (id, name, email, password, document, createdAt) VALUES ($1, $2, $3, $4, $5, $6)"

	_, err := r.db.Exec(sqlStatement, user.Id, user.Name, user.Email, user.Password, user.Document, user.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) GetByEmail(email string) (*canonical.User, error) {
	rows, err := r.db.Query(
		"SELECT * FROM Customer WHERE Email = $1",
		email,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var user canonical.User

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

	return &user, errorNotFound
}
