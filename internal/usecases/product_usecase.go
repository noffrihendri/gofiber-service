package usecases

import (
	"faber/internal/domain"
	"faber/internal/repositories"
	"log"

	fiberlog "github.com/gofiber/fiber/v2/log"
)

type ProductUsecase struct {
	//productRepo infrastructure.ProductRepoContract
	ProductRepository repositories.ProductRepositoryContract
	name              string
}

type ProductUsecaseContract interface {
	GetProduct() ([]*domain.Product, error)
	GetProductById(Id string) (*domain.Product, error)
}

func NewProductUsecase(productRepositoryContract repositories.ProductRepositoryContract) ProductUsecaseContract {
	return &ProductUsecase{
		ProductRepository: productRepositoryContract,
		name:              "Product Usecase",
	}

}

func (r ProductUsecase) GetProduct() ([]*domain.Product, error) {
	fiberlog.Info("Hello, World!")
	result, err := r.ProductRepository.ListProduct()
	if err != nil {
		log.Printf("[%s][Read] is executed\n", err)
		return nil, err
	}
	return result, nil
}
func (r ProductUsecase) GetProductById(Id string) (*domain.Product, error) {
	result, err := r.ProductRepository.GetProductById(Id)
	if err != nil {
		log.Printf("[%s][Read] is executed\n", err)
		return nil, err
	}
	return result, nil
}
