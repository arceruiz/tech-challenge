package repository

import (
	"database/sql"
	"errors"
	"tech-challenge/internal/config"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var (
	cfg           = &config.Cfg
	ErrorNotFound = errors.New("entity not found")
)

func New() *sql.DB {
	db, err := sql.Open("postgres", cfg.DB.ConnectionString)
	if err != nil {
		logrus.Fatal(err)
	}

	return db
}
