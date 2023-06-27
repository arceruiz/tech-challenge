package rest

import (
	"client/internal/middlewares"
	"client/internal/rest/port"
	"client/internal/service"

	"github.com/labstack/echo"
)

type rest struct {
	user    *port.UserPort
	product *port.ProductPort
}

func New() *rest {
	return &rest{
		user:    port.NewUserPort(service.NewUserService()),
		product: port.NewProductPort(service.NewProductService()),
	}
}

func (r *rest) Start() error {
	router := echo.New()

	router.Use(middlewares.Logger)

	mainPort := router.Group("/api")

	userPort := mainPort.Group("/user")
	r.user.Register(userPort)

	productPort := mainPort.Group("/product")
	r.product.Register(productPort)

	return router.Start(":3001")
}
