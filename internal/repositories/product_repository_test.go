package repositories

import (
	"faber/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	mockRepo := new(MockProductRepo)
	product := &domain.Product{
		Name:     "Test Product",
		Price:    10.0,
		Quantity: 100,
	}

	mockRepo.On("CreateProduct", product).Return(product, nil)

	result, err := mockRepo.CreateProduct(product)

	assert.NoError(t, err)
	assert.Equal(t, product, result)
	mockRepo.AssertExpectations(t)
}

func TestGetProduct(t *testing.T) {
	mockRepo := new(MockProductRepo)
	product := &domain.Product{
		ID:       "1",
		Name:     "Test Product",
		Price:    10.0,
		Quantity: 100,
	}

	mockRepo.On("GetProduct", "1").Return(product, nil)

	result, err := mockRepo.GetProductById("1")

	assert.NoError(t, err)
	assert.Equal(t, product, result)
	mockRepo.AssertExpectations(t)
}

func TestListProduct(t *testing.T) {
	mockRepo := new(MockProductRepo)
	products := []*domain.Product{
		{Name: "Product 1", Price: 10.0, Quantity: 100},
		{Name: "Product 2", Price: 20.0, Quantity: 200},
	}

	mockRepo.On("ListProduct").Return(products, nil)

	result, err := mockRepo.ListProduct()

	assert.NoError(t, err)
	assert.Equal(t, products, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateProduct(t *testing.T) {
	mockRepo := new(MockProductRepo)
	product := &domain.Product{
		ID:       "1",
		Name:     "Updated Product",
		Price:    20.0,
		Quantity: 200,
	}

	mockRepo.On("UpdateProduct", product).Return(nil)

	err := mockRepo.UpdateProduct(product)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteProduct(t *testing.T) {
	mockRepo := new(MockProductRepo)
	productID := "1"

	mockRepo.On("DeleteProduct", productID).Return(nil)

	err := mockRepo.DeleteProduct(productID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
