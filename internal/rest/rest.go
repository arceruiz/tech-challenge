package rest

import (
	"tech-challenge/internal/config"
	"tech-challenge/internal/middlewares"
	adapterRest "tech-challenge/internal/rest/adapter"
	"tech-challenge/internal/rest/port"

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
