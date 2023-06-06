package product

import (
	"stock-service/model"
	"stock-service/payload"
	"stock-service/repository"
	"stock-service/repository/product"
)

type ProductUseCase struct {
	repo product.IRepoProduct
}
type IProductUseCase interface {
	CreateProduct(req payload.CreateProductReq) (*model.Product, error)
	GetProductById(id int) (*model.Product, error)
	GetListProduct() ([]model.Product, error)
	UpdateQuantity(req payload.UpdateQuantity) error
}

func New(repo *repository.Repository) IProductUseCase {
	return &ProductUseCase{
		repo: repo.RepoProduct,
	}
}
