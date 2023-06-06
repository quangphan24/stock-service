package usecase

import (
	"stock-service/repository"
	"stock-service/usecase/product"
)

type UseCase struct {
	ProductUseCase product.IProductUseCase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{
		ProductUseCase: product.New(repo),
	}
}
