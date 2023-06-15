package api

import (
	controllers "tech-challenge/internal/API/Controllers"
	"tech-challenge/internal/API/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type rest struct {
	customerController controllers.CustomerController
	productController  controllers.ProductController
	orderController    controllers.OrderController
	paymentController  controllers.PaymentController
}

func New() Channel {
	return &rest{}
}

func (r *rest) Start() error {
	e := echo.New()
	e.Use(middlewares.NewLogger())
	e.Use(middleware.Recover())

	// Customers
	e.GET("/customers", r.customerController.GetCustomers)
	e.POST("/customers", r.customerController.CreateCustomer)
	e.GET("/customers/:id", r.customerController.GetCustomer)
	e.PUT("/customers/:id", r.customerController.UpdateCustomer)
	e.DELETE("/customers/:id", r.customerController.DeleteCustomer)

	// Products
	e.GET("/products", r.productController.GetProducts)
	e.POST("/products", r.productController.CreateProduct)
	e.GET("/products/:id", r.productController.GetProduct)
	e.PUT("/products/:id", r.productController.UpdateProduct)
	e.DELETE("/products/:id", r.productController.DeleteProduct)

	// Orders
	e.GET("/orders", r.orderController.GetOrders)
	e.POST("/orders", r.orderController.CreateOrder)
	e.GET("/orders/:id", r.orderController.GetOrder)
	e.PUT("/orders/:id", r.orderController.UpdateOrder)
	e.DELETE("/orders/:id", r.orderController.DeleteOrder)

	// Payments
	e.GET("/payments", r.paymentController.GetPayments)
	e.POST("/payments", r.paymentController.CreatePayment)
	e.GET("/payments/:id", r.paymentController.GetPayment)
	e.PUT("/payments/:id", r.paymentController.UpdatePayment)
	e.DELETE("/payments/:id", r.paymentController.DeletePayment)

	return e.Start(":8080") //+ config.Instance.Server.HttpPort)
}
