package repositories

import (
	"faber/internal/domain"

	"github.com/stretchr/testify/mock"
)

type MockProductRepo struct {
	mock.Mock
}

func (m *MockProductRepo) CreateProduct(product *domain.Product) (*domain.Product, error) {
	args := m.Called(product)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *MockProductRepo) UpdateProduct(product *domain.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepo) ListProduct() ([]*domain.Product, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Product), args.Error(1)
}

func (m *MockProductRepo) GetProductById(ID string) (*domain.Product, error) {
	args := m.Called(ID)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *MockProductRepo) DeleteProduct(ID string) error {
	args := m.Called(ID)
	return args.Error(0)
}
