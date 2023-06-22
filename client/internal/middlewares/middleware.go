package middlewares

import (
	"client/internal/auth/token"
	"context"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func Logger(fx echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		log := logrus.WithContext(context.Background())
		request := ctx.Request()
		log.WithFields(logrus.Fields{
			log.Message: request.RequestURI,
			log.Message: request.Method,
			log.Message: request.Host,
		})
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