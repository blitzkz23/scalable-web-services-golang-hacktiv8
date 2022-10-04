package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarID string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var CarDatas = []Car{}

// * Fungsi2 di bawah ini akan menjadi endpoint untuk proses create data dengan parameter *gin.Context.  *gin.Context mempunyai berbagai macam method untuk mendapat request body dari client, mengirim response, dll.
// * Create data menggunakan JSON raw response pada postman, contoh:
/**
	{
		"brand": "BMW",
		"price": 21000000,
		"model": "Super"
	}
**/
// context dapat menggantika request dan response writer pada http
func CreateCar(ctx *gin.Context) {
	var newCar Car

	// * Method ShouldBindJSON merupakan method *gin.Context yang digunakan untuk mem-binding data JSON yang dikirimkan oleh client sebagai request body kepada server
	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		// * Method AbortWithError digunakan untuk melempar error jika ada.  Parameter pertama adalah status code, parameter kedua adalah error yang akan dilempar.
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newCar.CarID = fmt.Sprintf("c%d", len(CarDatas)+1)
	CarDatas = append(CarDatas, newCar)

	// * Mengirim data response kepada client dengan JSON.  Parameter pertama adalah status code, parameter kedua adalah data yang akan dikirimkan.
	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
	})
}

// * Update car
func UpdateCar(ctx *gin.Context) {
	// Menerima params dari router
	carID := ctx.Param("carID")
	condition := false

	var updatedCar Car

	if err := ctx.ShouldBindJSON(&updatedCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Mencari data yang akan diupdate dengan looping
	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			CarDatas[i] = updatedCar
			CarDatas[i].CarID = carID
			break
		}
	}

	// Jika data tidak ditemukan, error status 404
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been updated", carID),
	})
}

// * Get car
func GetCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var carData Car

	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			carData = CarDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car": carData,
	})
}

// * Delete car
func DeleteCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var carIndex int

	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			carIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	copy(CarDatas[carIndex:], CarDatas[carIndex+1:])
	CarDatas[len(CarDatas)-1] = Car{}
	CarDatas = CarDatas[:len(CarDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been sucessfully deleted", carID),
	})
}
