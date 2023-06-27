package repository

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var (
	ErrorNotFound = errors.New("entity not found")
)

func New() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=5432 dbname=fiap_tech_challenge user=postgres password=1234 sslmode=disable")
	if err != nil {
		logrus.Fatal(err)
	}

	return db
}
