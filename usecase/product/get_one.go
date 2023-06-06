package product

import "stock-service/model"

func (uc *ProductUseCase) GetProductById(id int) (*model.Product, error) {
	return uc.repo.GetProductById(id)
}
