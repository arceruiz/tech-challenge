package rest

import (
	"fmt"
	"net/http"
	"tech-challenge/internal/canonical"
	"tech-challenge/internal/service"

	"github.com/labstack/echo"
)

type Order interface {
	RegisterGroup(g *echo.Group)
	Get(c echo.Context) error
	CheckoutOrder(c echo.Context) error
}

type order struct {
	service service.OrderService
}

func NewOrderChannel() Order {
	return &order{
		service: service.NewOrderService(),
	}
}

func (p *order) RegisterGroup(g *echo.Group) {
	g.GET("/", p.Get)
	g.POST("/checkout", p.CheckoutOrder)
	g.POST("/create", p.Create)
}

func (p *order) Get(c echo.Context) error {
	orderID := c.QueryParam("id")
	status := c.QueryParam("status")

	if orderID != "" {
		order, err := p.service.GetByID(c.Request().Context(), orderID)
		if err != nil {
			return c.JSON(http.StatusNotFound, "Order not found")
		}
		return c.JSON(http.StatusOK, order)
	}

	if status != "" {
		orders, err := p.service.GetByStatus(c.Request().Context(), status)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, orders)
	}

	orders, err := p.service.GetOrders(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, orders)
}

func (p *order) Create(c echo.Context) error {
	var orderRequest OrderRequest

	if err := c.Bind(&orderRequest); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Message: fmt.Errorf("invalid data").Error(),
		})
	}

	id, err := p.service.CreateOrder(c.Request().Context(), orderRequest.toCanonical())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, id)
}

func (p *order) CheckoutOrder(c echo.Context) error {
	orderID := c.Param("id")

	var payment canonical.Payment
	err := c.Bind(&payment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	err = p.service.CheckoutOrder(c.Request().Context(), orderID, payment)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	return c.JSON(http.StatusOK, nil)
}
