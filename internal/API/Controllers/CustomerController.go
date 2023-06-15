package controllers

import (
	"net/http"
	usecases "tech-challenge/internal/Application/UseCases"
	entities "tech-challenge/internal/Domain/Entities"

	"github.com/labstack/echo"
)

type CustomerController struct {
	service usecases.ICustomerUseCase
}

func NewCustomerController(service usecases.ICustomerUseCase) *CustomerController {
	return &CustomerController{
		service: service,
	}
}

func (h *CustomerController) GetCustomers(c echo.Context) error {
	customers, err := h.service.GetCustomers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, customers)
}

func (h *CustomerController) CreateCustomer(c echo.Context) error {
	var customer entities.Customer
	if err := c.Bind(&customer); err != nil {
		return err
	}
	createdCustomer, err := h.service.CreateCustomer(customer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusCreated, createdCustomer)
}

func (h *CustomerController) GetCustomer(c echo.Context) error {
	id := c.Param("id")
	customer, err := h.service.GetCustomer(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, customer)
}

func (h *CustomerController) UpdateCustomer(c echo.Context) error {
	id := c.Param("id")
	var updatedCustomer entities.Customer
	if err := c.Bind(&updatedCustomer); err != nil {
		return err
	}
	updatedCustomer, err := h.service.UpdateCustomer(id, updatedCustomer)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, updatedCustomer)
}

func (h *CustomerController) DeleteCustomer(c echo.Context) error {
	id := c.Param("id")
	err := h.service.DeleteCustomer(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.NoContent(http.StatusNoContent)
}
