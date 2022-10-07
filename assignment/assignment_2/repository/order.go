package repository

import "scalable-web-services-golang-hacktiv8/assignment/assignment_2/entity"

// ! Repository interface dari order
type OrderRepository interface {
	InsertOrder(order *entity.Order) (*entity.Order, error)
	GetAllOrders() ([]*entity.Order, error)
}
