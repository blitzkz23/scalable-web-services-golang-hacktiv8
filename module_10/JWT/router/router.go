package router

import (
	"scalable-web-services-golang-hacktiv8/module_10/JWT/controllers"
	"scalable-web-services-golang-hacktiv8/module_10/JWT/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		// Implementasi middleware JWT
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:productID", middlewares.ProductAuthorization(), controllers.UpdateProduct)
	}

	return r
}
