package service

import (
	"scalable-web-services-golang-hacktiv8/assignment/assignment_2/dto"
	"scalable-web-services-golang-hacktiv8/assignment/assignment_2/entity"
	"scalable-web-services-golang-hacktiv8/assignment/assignment_2/repository"
)

// ! OrderService interface
type OrderService interface {
	InsertOrder(*dto.NewOrderRequest) (*dto.NewOrderResponse, error)
	GetAllOrders() ([]*entity.Order, error)
}

// ! OrderService implementation
type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{
		repo: repo,
	}
}

func (o *orderService) InsertOrder(orderPayload *dto.NewOrderRequest) (*dto.NewOrderResponse, error) {
	// ! Service untuk insert data order ke database
	if err := orderPayload.Validate(); err != nil {
		return nil, err
	}

	orderRequest := &entity.Order{
		Customer_name: orderPayload.Customer_name,
	}

	newOrder, err := o.repo.InsertOrder(orderRequest)
	if err != nil {
		return nil, err
	}

	return newOrder.NewOrderResponseDTO(), nil

}

func (o *orderService) GetAllOrders() ([]*entity.Order, error) {
	// ! Service untuk data order dari database
	orders, err := o.repo.GetAllOrders()
	if err != nil {
		return nil, err
	}

	return orders, nil
}
