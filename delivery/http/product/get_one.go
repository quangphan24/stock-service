package product

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (r *Route) GetProductById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	newProduct, err := r.useCase.ProductUseCase.GetProductById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, newProduct)
}
