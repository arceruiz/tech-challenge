package controllers

import (
	"net/http"
	usecases "tech-challenge/internal/Application/UseCases"
	entities "tech-challenge/internal/Domain/Entities"

	"github.com/labstack/echo"
)

type OrderController struct {
	service usecases.IOrderUseCase
}

func NewOrderController(service usecases.IOrderUseCase) *OrderController {
	return &OrderController{
		service: service,
	}
}

func (h *OrderController) GetOrders(c echo.Context) error {
	orders, err := h.service.GetOrders()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, orders)
}

func (h *OrderController) CreateOrder(c echo.Context) error {
	var order entities.Order
	if err := c.Bind(&order); err != nil {
		return err
	}
	createdOrder, err := h.service.CreateOrder(order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusCreated, createdOrder)
}

func (h *OrderController) GetOrder(c echo.Context) error {
	id := c.Param("id")
	order, err := h.service.GetOrder(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, order)
}

func (h *OrderController) UpdateOrder(c echo.Context) error {
	id := c.Param("id")
	var updatedOrder entities.Order
	if err := c.Bind(&updatedOrder); err != nil {
		return err
	}
	updatedOrder, err := h.service.UpdateOrder(id, updatedOrder)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, updatedOrder)
}

func (h *OrderController) DeleteOrder(c echo.Context) error {
	id := c.Param("id")
	err := h.service.DeleteOrder(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.NoContent(http.StatusNoContent)
}
