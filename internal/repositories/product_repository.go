package repositories

import (
	"faber/internal/domain"
	"fmt"
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
	CreateProduct(product *domain.Product) (*domain.Product, error)
	UpdateProduct(product *domain.Product) error
	DeleteProduct(ID string) error
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

func (r *ProductRepository) CreateProduct(product *domain.Product) (*domain.Product, error) {
	log.Printf("[%s][Create] is executed\n", r.name)

	sql := fmt.Sprintf("insert into product(name,price,quantity) values('%s',%f,%d) RETURNING id,name,price,quantity", product.Name, product.Price, product.Quantity)

	if err := r.db.Raw(sql).Scan(&product).Error; err != nil {
		log.Printf("Error : [%s][Create] %s\n", r.name, err.Error())
		return product, nil
	}
	fmt.Println("model created", product)
	return product, nil
}
func (r *ProductRepository) UpdateProduct(product *domain.Product) error {
	log.Printf("[%s][Update] is executed\n", r.name)

	if err := r.db.Model(&product).Updates(&product).Error; err != nil {
		log.Printf("Error : [%s][Update] %s", r.name, err.Error())
		return err
	}

	return nil
}

func (r *ProductRepository) DeleteProduct(ID string) error {
	log.Printf("[%s][Delete] is executed\n", r.name)

	var product domain.Product

	if err := r.db.Table("product").Delete(&product, map[string]string{"id": ID}).Error; err != nil {
		log.Printf("Error : [%s][Delete] %s", r.name, err.Error())
		return err
	}

	return nil
}
