package product

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"stock-service/model"
	"stock-service/payload"
)

type RepoProduct struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) IRepoProduct {
	return &RepoProduct{db: db}
}

//go:generate mockery --name IRepoProduct
type IRepoProduct interface {
	CreateProduct(product *model.Product) error
	GetProductById(id int) (*model.Product, error)
	GetListProduct() ([]model.Product, error)
	UpdateQuantity(req payload.UpdateQuantity) error
}

func (r *RepoProduct) CreateProduct(product *model.Product) error {
	tx := r.db
	return tx.Table("product").Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "email"}}, UpdateAll: true}).
		Create(&product).Error
}
func (r *RepoProduct) GetProductById(id int) (*model.Product, error) {
	product := &model.Product{}

	err := r.db.Table("product").Where("id=?", id).Take(&product).Error
	return product, err
}
func (r *RepoProduct) GetListProduct() ([]model.Product, error) {
	products := []model.Product{}

	err := r.db.Table("product").Find(&products).Error
	return products, err
}
func (r *RepoProduct) UpdateQuantity(req payload.UpdateQuantity) error {
	err := r.db.Table("product").Where("id = ?", req.Id).Updates(map[string]interface{}{"quantity": req.Quantity}).Error
	return err
}
