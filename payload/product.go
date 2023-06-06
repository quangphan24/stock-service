package payload

import (
	"github.com/go-playground/validator/v10"
)

type CreateProductReq struct {
	Name     string `json:"name" validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
	Price    int    `json:"price" validate:"required"`
}

func (r *CreateProductReq) Validate(v validator.Validate) error {
	if err := v.Struct(r); err != nil {
		return err
	}
	return nil
}

type UpdateQuantity struct {
	Id       int `json:"id" validate:"required"`
	Quantity int `json:"quantity" validate:"required"`
}

func (r *UpdateQuantity) Validate(v validator.Validate) error {
	if err := v.Struct(r); err != nil {
		return err
	}
	return nil
}
