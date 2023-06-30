package middlewares

import (
	"context"
	"net/http"
	"tech-challenge/internal/auth/token"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func Logger(fx echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		log := logrus.WithContext(context.Background())
		request := ctx.Request()
		log.WithFields(logrus.Fields{
			"Host":   request.Host,
			"URI":    request.RequestURI,
			"Method": request.Method,
		}).Info()

		return fx(ctx)
	}
}

func Authorization(fx echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if err := token.ValidateToken(ctx.Request()); err != nil {
			ctx.Response().Header().Set("Content-Type", "application/json")
			ctx.Response().WriteHeader(http.StatusUnauthorized)
			return err
		}

		return fx(ctx)
	}
}
