package dto

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type NewOrderRequest struct {
	Customer_name string `json:"customer_name" valid:"required~customer_name cannot be empty"`
}

func (o *NewOrderRequest) Validate() error {
	_, err := govalidator.ValidateStruct(o)
	if err != nil {
		return err
	}
	return nil
}

type NewOrderResponse struct {
	ID            uint      `json:"id"`
	Customer_name string    `json:"customer_name"`
	Ordered_at    time.Time `json:"ordered_at"`
}
