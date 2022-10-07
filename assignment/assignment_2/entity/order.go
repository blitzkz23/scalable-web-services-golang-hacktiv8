package entity

import (
	"scalable-web-services-golang-hacktiv8/assignment/assignment_2/dto"
	"time"
)

type Order struct {
	ID            uint      `json:"id"`
	Customer_name string    `json:"customer_name"`
	Ordered_at    time.Time `json:"ordered_at"`
	Items         []Item    `json:"items"`
}

func (o *Order) NewOrderResponseDTO() *dto.NewOrderResponse {
	return &dto.NewOrderResponse{
		ID:            o.ID,
		Customer_name: o.Customer_name,
		Ordered_at:    o.Ordered_at,
	}
}
