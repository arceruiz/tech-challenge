package adapter

import (
	"net/http"
	"tech-challenge/internal/canonical"
	rest "tech-challenge/internal/rest/port"
	"tech-challenge/internal/service/adapter"
	services "tech-challenge/internal/service/port"

	"github.com/labstack/echo"
)

// GetOrders() ([]canonical.Order, error)
// CreateOrder(order canonical.Order) error
// UpdateOrder(id string, updatedOrder canonical.Order) (canonical.Order, error)
// GetByID(id string) (*canonical.Order, error)
// GetByStatus(id string) ([]canonical.Order, error)
// CheckoutOrder(orderID string, payment canonical.Payment) error
type OrderPort struct {
	service services.OrderService
}

func NewOrderPort() rest.OrderPort {
	return &OrderPort{
		service: adapter.NewOrderService(),
	}
}

func (p *OrderPort) Register(g *echo.Group) {
	indexPath := ""
	g.GET(indexPath, p.Get)
	g.POST(indexPath+"/checkout", p.CheckoutOrder)
}

func (p *OrderPort) Get(c echo.Context) error {
	productID := c.QueryParam("id")
	status := c.QueryParam("status")

	if productID != "" {
		product, err := p.service.GetByID(productID)
		if err != nil {
			return c.JSON(http.StatusNotFound, "Order not found")
		}
		return c.JSON(http.StatusOK, product)
	}

	if status != "" {
		products, err := p.service.GetByStatus(status)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, products)
	}

	products, err := p.service.GetOrders()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, products)
}

func (p *OrderPort) CheckoutOrder(c echo.Context) error {
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
