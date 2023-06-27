package repository

import (
	"errors"

	_ "github.com/lib/pq"
)

var (
	errorNotFound = errors.New("entity not found")
)
