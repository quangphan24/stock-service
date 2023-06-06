package product

import (
	"github.com/labstack/echo/v4"
	"stock-service/usecase"
)

type Route struct {
	useCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{useCase: useCase}

	group.POST("", r.Create)
	group.GET("/:id", r.GetProductById)
	group.GET("", r.GetListProduct)
}
