package rest

import (
	"client/internal/canonical"
	"errors"
	"net/http"
)

func HandleError(err error) int {
	if errors.Is(err, canonical.ErrorNotFound) {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
