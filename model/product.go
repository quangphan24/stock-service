package model

type Product struct {
	BaseModel
	ID       int    `json:"id"`
	Name     string `json:"name" validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
	Price    int    `json:"price" validate:"required"`
}
