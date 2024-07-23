package repositories

import (
	"faber/internal/domain"
	"log"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db   *gorm.DB
	name string
}

type ProductRepositoryContract interface {
	ListProduct() ([]*domain.Product, error)
	GetProductById(Id string) (*domain.Product, error)
}

func NewProductRepository(db *gorm.DB) ProductRepositoryContract {
	return &ProductRepository{
		db:   db,
		name: "Product Repository",
	}
}

func (r *ProductRepository) ListProduct() ([]*domain.Product, error) {
	log.Printf("[%s][List] is executed\n", r.name)

	var products []*domain.Product

	db := r.db

	db.Table("product").Find(&products)

	return products, nil
}
func (r *ProductRepository) GetProductById(Id string) (*domain.Product, error) {
	log.Printf("[%s][Get] is executed\n", r.name)

	var product domain.Product

	if err := r.db.Table("product").Where("id = ?", Id).First(&product).Error; err != nil {
		log.Printf("Error : [%s][GET] %s", r.name, err.Error())
		return &product, err
	}

	return &product, nil

}
