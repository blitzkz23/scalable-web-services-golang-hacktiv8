package routers

import (
	"scalable-web-services-golang-hacktiv8/module_6/gin_framework/controllers"

	"github.com/gin-gonic/gin"
)

// * Fungsi ini akan mengembalikan *gin.Engine yang merupakan router dari gin untuk menjalankan server dari aplikasi.
func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)
	router.PUT("/cars/:carID", controllers.UpdateCar)
	router.GET("/cars/:carID", controllers.GetCar)
	router.DELETE("/cars/:carID", controllers.DeleteCar)

	// bisa buat route group
	// carsRoute := router.Group{
	// 	router.POST("/", controllers.CreateCar),
	// 	router.PUT("/", controllers.UpdateCar),
	// 	router.GET("/", controllers.GetCar),
	// 	router.DELETE("/", controllers.DeleteCar),
	// }

	return router
}
