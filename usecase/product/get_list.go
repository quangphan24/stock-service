package product

import "stock-service/model"

func (u *ProductUseCase) GetListProduct() ([]model.Product, error) {
	return u.repo.GetListProduct()
}
