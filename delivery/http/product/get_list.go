package product

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (r *Route) GetListProduct(c echo.Context) error {
	products, err := r.useCase.ProductUseCase.GetListProduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, products)
}
