package rest

import (
	"client/internal/config"
	"client/internal/middlewares"
	adapterRest "client/internal/rest/adapter"
	"client/internal/rest/port"

	"github.com/labstack/echo"
)

var (
	cfg = &config.Cfg
)

type rest struct {
	customer port.CustomerPort
	product  port.ProductPort
}

func New() rest {
	return rest{
		customer: adapterRest.NewCustomerPort(),
		product:  adapterRest.NewProductPort(),
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

	return router.Start(":" + cfg.Server.Port)
}
