package product

import (
	"github.com/sirupsen/logrus"
	"stock-service/model"
	"stock-service/payload"
)

func (uc *ProductUseCase) CreateProduct(req payload.CreateProductReq) (*model.Product, error) {

	product := &model.Product{
		Name:     req.Name,
		Quantity: req.Quantity,
		Price:    req.Price,
	}
	//create
	if err := uc.repo.CreateProduct(product); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return product, nil
}
