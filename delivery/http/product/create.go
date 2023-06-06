package product

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"stock-service/payload"
)

func (r *Route) Create(c echo.Context) error {
	var req payload.CreateProductReq

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	validate := validator.New()
	if err := req.Validate(*validate); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	newProduct, err := r.useCase.ProductUseCase.CreateProduct(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, newProduct)
}
