package repository

import (
	"gorm.io/gorm"
	"stock-service/repository/product"
)

type Repository struct {
	RepoProduct product.IRepoProduct
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		RepoProduct: product.NewRepo(db),
	}
}
