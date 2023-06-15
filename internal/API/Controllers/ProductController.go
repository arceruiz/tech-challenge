package controllers

import (
	"net/http"
	usecases "tech-challenge/internal/Application/UseCases"
	entities "tech-challenge/internal/Domain/Entities"

	"github.com/labstack/echo"
)

type ProductController struct {
	service usecases.IProductUseCase
}

func NewProductController(service usecases.IProductUseCase) *ProductController {
	return &ProductController{
		service: service,
	}
}

func (h *ProductController) GetProducts(c echo.Context) error {
	products, err := h.service.GetProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, products)
}

func (h *ProductController) CreateProduct(c echo.Context) error {
	var product entities.Product
	if err := c.Bind(&product); err != nil {
		return err
	}
	createdProduct, err := h.service.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusCreated, createdProduct)
}

func (h *ProductController) GetProduct(c echo.Context) error {
	id := c.Param("id")
	product, err := h.service.GetProduct(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, product)
}

func (h *ProductController) UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	var updatedProduct entities.Product
	if err := c.Bind(&updatedProduct); err != nil {
		return err
	}
	updatedProduct, err := h.service.UpdateProduct(id, updatedProduct)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, updatedProduct)
}

func (h *ProductController) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	err := h.service.DeleteProduct(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.NoContent(http.StatusNoContent)
}
