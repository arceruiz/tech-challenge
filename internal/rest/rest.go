package rest

import (
	"tech-challenge/internal/config"
	"tech-challenge/internal/middlewares"

	"github.com/labstack/echo"
)

var (
	cfg = &config.Cfg
)

type rest struct {
	customer Customer
	product  Product
	order    Order
}

func New() rest {
	return rest{
		customer: NewCustomerChannel(),
		product:  NewProductChannel(),
		order:    NewOrderChannel(),
	}
}

func (r rest) Start() error {
	router := echo.New()

	router.Use(middlewares.Logger)

	mainGroup := router.Group("/api")

	customerGroup := mainGroup.Group("/user")
	r.customer.Register(customerGroup)

	productGroup := mainGroup.Group("/product")
	r.product.Register(productGroup)
	productGroup.Use(middlewares.Authorization)

	orderGroup := mainGroup.Group("/order")
	r.order.Register(orderGroup)
	productGroup.Use(middlewares.Authorization)

	return router.Start(":" + cfg.Server.Port)
}
