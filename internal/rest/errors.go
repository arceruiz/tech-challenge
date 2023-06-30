package rest

import (
	"errors"
	"net/http"
	"tech-challenge/internal/canonical"
)

func HandleError(err error) int {
	if errors.Is(err, canonical.ErrorNotFound) {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
