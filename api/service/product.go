package service

import (
	"InjectGo-Workshop/api/repository"
	"InjectGo-Workshop/model"
)

type ProductService interface {
	CreateProduct(user *model.Product) (*model.Product, error)
}

type IProductService struct {
	UserRepository    repository.UserRepository
	ProductRepository repository.ProductRepository
}

func NewProductService(UserRepository repository.UserRepository, ProductRepository repository.ProductRepository) ProductService {
	return &IProductService{
		UserRepository:    UserRepository,
		ProductRepository: ProductRepository,
	}
}

func (s *IProductService) CreateProduct(product *model.Product) (*model.Product, error) {
	return s.CreateProduct(product)
}
