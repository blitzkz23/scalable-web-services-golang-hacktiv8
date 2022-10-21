package service

import (
	"errors"
	"scalable-web-services-golang-hacktiv8/module_11/mock_test/entity"
	"scalable-web-services-golang-hacktiv8/module_11/mock_test/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id string) (*entity.Product, error) {
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}
