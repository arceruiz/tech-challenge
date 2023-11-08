package rest

import (
	"tech-challenge/internal/config"
	"tech-challenge/internal/middlewares"

	"github.com/labstack/echo/v4"
)

var (
	cfg = &config.Cfg
)

type rest struct {
	customer Customer
	product  Product
	order    Order
	payment  Payment
}

func New() rest {
	return rest{
		customer: NewCustomerChannel(),
		product:  NewProductChannel(),
		order:    NewOrderChannel(),
		payment:  NewPaymentChannel(),
	}
}

func (r rest) Start() error {
	router := echo.New()

	router.Use(middlewares.Logger)

	mainGroup := router.Group("/api")

	customerGroup := mainGroup.Group("/user")
	r.customer.RegisterGroup(customerGroup)

	productGroup := mainGroup.Group("/product")
	r.product.RegisterGroup(productGroup)
	productGroup.Use(middlewares.Authorization)

	orderGroup := mainGroup.Group("/order")
	r.order.RegisterGroup(orderGroup)
	orderGroup.Use(middlewares.Authorization)

	paymentGroup := mainGroup.Group("/payment")
	r.payment.RegisterGroup(paymentGroup)
	paymentGroup.Use(middlewares.Authorization)

	return router.Start(":" + cfg.Server.Port)
}
