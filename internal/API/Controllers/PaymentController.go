package controllers

import (
	"net/http"
	usecases "tech-challenge/internal/Application/UseCases"
	entities "tech-challenge/internal/Domain/Entities"

	"github.com/labstack/echo"
)

type PaymentController struct {
	service usecases.IPaymentUseCase
}

func NewPaymentController(service usecases.IPaymentUseCase) *PaymentController {
	return &PaymentController{
		service: service,
	}
}

func (h *PaymentController) GetPayments(c echo.Context) error {
	payments, err := h.service.GetPayments()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, payments)
}

func (h *PaymentController) CreatePayment(c echo.Context) error {
	var payment entities.Payment
	if err := c.Bind(&payment); err != nil {
		return err
	}
	createdPayment, err := h.service.CreatePayment(payment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusCreated, createdPayment)
}

func (h *PaymentController) GetPayment(c echo.Context) error {
	id := c.Param("id")
	payment, err := h.service.GetPayment(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, payment)
}

func (h *PaymentController) UpdatePayment(c echo.Context) error {
	id := c.Param("id")
	var updatedPayment entities.Payment
	if err := c.Bind(&updatedPayment); err != nil {
		return err
	}
	updatedPayment, err := h.service.UpdatePayment(id, updatedPayment)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, updatedPayment)
}

func (h *PaymentController) DeletePayment(c echo.Context) error {
	id := c.Param("id")
	err := h.service.DeletePayment(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.NoContent(http.StatusNoContent)
}
