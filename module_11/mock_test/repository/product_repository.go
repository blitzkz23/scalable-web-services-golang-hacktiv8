package repository

import "scalable-web-services-golang-hacktiv8/module_11/mock_test/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
}
