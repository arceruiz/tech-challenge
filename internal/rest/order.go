package rest

import (
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
	indexPath := ""
	g.GET(indexPath, p.Get)
	g.POST(indexPath+"/checkout", p.CheckoutOrder)
}

func (p *order) Get(c echo.Context) error {
	orderID := c.QueryParam("id")
	status := c.QueryParam("status")

	if orderID != "" {
		order, err := p.service.GetByID(orderID)
		if err != nil {
			return c.JSON(http.StatusNotFound, "Order not found")
		}
		return c.JSON(http.StatusOK, order)
	}

	if status != "" {
		orders, err := p.service.GetByStatus(status)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, orders)
	}

	orders, err := p.service.GetOrders()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, orders)
}

func (p *order) CheckoutOrder(c echo.Context) error {
	orderID := c.Param("id")

	var payment canonical.Payment
	err := c.Bind(&payment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	err = p.service.CheckoutOrder(orderID, payment)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	return c.JSON(http.StatusOK, nil)
}