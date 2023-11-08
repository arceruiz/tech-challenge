package rest

import (
	"context"
	"fmt"
	"net/http"
	"tech-challenge/internal/service"

	"github.com/labstack/echo/v4"
)

type Order interface {
	RegisterGroup(g *echo.Group)
	Get(c echo.Context) error
	CheckoutOrder(c echo.Context) error
	Update(c echo.Context) error
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
	g.GET("", p.GetAll)
	g.GET("/:id", p.Get)
	g.POST("/checkout", p.CheckoutOrder)
	g.POST("/", p.Create)
	//g.POST("/", p.ready) needs to be implemented
	g.PUT("/:id", p.Update)

}

func (p *order) GetAll(ctx echo.Context) error {
	status := ctx.QueryParam("status")

	response, err := p.get(ctx.Request().Context(), "", status)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, response)
}

func (p *order) Get(ctx echo.Context) error {
	id := ctx.Param("id")

	response, err := p.get(ctx.Request().Context(), id, "")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	if len(response) == 0 {
		return ctx.NoContent(http.StatusNotFound)
	}

	return ctx.JSON(http.StatusOK, response[0])
}

func (p *order) get(ctx context.Context, orderID string, status string) ([]OrderResponse, error) {
	if orderID != "" {
		order, err := p.service.GetByID(ctx, orderID)
		if err != nil {
			return nil, err
		}

		if order == nil {
			return nil, nil
		}

		return []OrderResponse{orderToResponse(*order)}, nil
	}

	var response []OrderResponse
	if status != "" {
		orders, err := p.service.GetByStatus(ctx, status)
		if err != nil {
			return nil, err
		}

		for _, order := range orders {
			response = append(response, orderToResponse(order))
		}

		return response, nil
	}

	orders, err := p.service.GetOrders(ctx)
	if err != nil {
		return nil, err
	}

	for _, order := range orders {
		response = append(response, orderToResponse(order))
	}

	return response, nil
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

	return c.JSON(http.StatusOK, OrderResponse{
		ID: id,
	})
}

func (p *order) Update(c echo.Context) error {
	orderID := c.Param("id")
	if len(orderID) == 0 {
		return c.JSON(http.StatusBadRequest, Response{
			Message: "missing id query param",
		})
	}

	var orderRequest OrderRequest
	if err := c.Bind(&orderRequest); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Message: fmt.Errorf("invalid data").Error(),
		})
	}

	err := p.service.UpdateOrder(c.Request().Context(), orderID, orderRequest.toCanonical())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Message: err.Error(),
		})
	}

	return c.NoContent(http.StatusOK)
}

func (p *order) CheckoutOrder(c echo.Context) error {
	orderID := c.QueryParam("id")
	if len(orderID) == 0 {
		return c.JSON(http.StatusBadRequest, Response{
			Message: "missing id query param",
		})
	}

	var payment PaymentRest
	err := c.Bind(&payment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	order, err := p.service.CheckoutOrder(c.Request().Context(), orderID, payment.toCanonical())
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	return c.JSON(http.StatusOK, orderToResponse(*order))
}
