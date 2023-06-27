package rest

import (
	"client/internal/middlewares"
	"client/internal/rest/port"
	"client/internal/service/adapter"

	"github.com/labstack/echo"
)

type rest struct {
	customer *port.CustomerPort
	product  *port.ProductPort
}

func New() *rest {
	return &rest{
		customer: port.NewUserPort(adapter.NewCustomerService()),
		product:  port.NewProductPort(adapter.NewProductService()),
	}
}

func (r *rest) Start() error {
	router := echo.New()

	router.Use(middlewares.Logger)

	mainPort := router.Group("/api")

	customerPort := mainPort.Group("/user")
	r.customer.Create(customerPort)

	productPort := mainPort.Group("/product")
	r.product.Register(productPort)

	return router.Start(":3001")
}
