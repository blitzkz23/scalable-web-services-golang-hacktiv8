package rest

import (
	"fmt"
	"scalable-web-services-golang-hacktiv8/assignment/assignment_2/database"
	"scalable-web-services-golang-hacktiv8/assignment/assignment_2/repository/pg"
	"scalable-web-services-golang-hacktiv8/assignment/assignment_2/service"

	"github.com/gin-gonic/gin"
)

const port = ":8080"

func StartApp() {
	database.InitializeDB()

	db := database.GetDB()

	// ! Inject Dependencies
	orderRepo := pg.NewOrderPg(db)
	orderService := service.NewOrderService(orderRepo)
	orderRestHandler := NewOrderRestHandler(orderService)

	// ! Routing API dengan GIN
	route := gin.Default()

	orderRoute := route.Group("/orders")
	{
		orderRoute.GET("/", orderRestHandler.getAllOrders)
		orderRoute.POST("/", orderRestHandler.insertOrder)
	}

	fmt.Println("Server is running on port", port)
	route.Run(port)
}
