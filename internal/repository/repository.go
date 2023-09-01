package repository

import (
	"context"
	"errors"
	"tech-challenge/internal/config"

	"github.com/jackc/pgx/v4"

	"github.com/sirupsen/logrus"
)

var (
	cfg           = &config.Cfg
	ErrorNotFound = errors.New("entity not found")
)

func New() *pgx.Conn {
	db, err := pgx.Connect(context.Background(), cfg.DB.ConnectionString)
	if err != nil {
		logrus.Fatal(err)
	}

	return db
}
