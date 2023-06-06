package product

import (
	"github.com/sirupsen/logrus"
	"stock-service/payload"
)

func (u *ProductUseCase) UpdateQuantity(req payload.UpdateQuantity) error {
	product, err := u.GetProductById(req.Id)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	req.Quantity = product.Quantity - req.Quantity
	return u.repo.UpdateQuantity(req)
}
