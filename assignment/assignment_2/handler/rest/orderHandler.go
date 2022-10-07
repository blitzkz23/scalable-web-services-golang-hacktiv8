package rest

import (
	"net/http"
	"scalable-web-services-golang-hacktiv8/assignment/assignment_2/dto"
	"scalable-web-services-golang-hacktiv8/assignment/assignment_2/service"

	"github.com/gin-gonic/gin"
)

type orderRestHandler struct {
	service service.OrderService
}

func NewOrderRestHandler(service service.OrderService) orderRestHandler {
	return orderRestHandler{
		service: service,
	}
}

func (o orderRestHandler) getAllOrders(c *gin.Context) {
	orders, err := o.service.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (o orderRestHandler) insertOrder(c *gin.Context) {
	var order dto.NewOrderRequest

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"message": "Invalid JSON Request",
			"err":     err.Error(),
		})
		return
	}

	newOrder, err := o.service.InsertOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"err":     "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusCreated, newOrder)
}
