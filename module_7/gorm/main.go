package main

import (
	"errors"
	"fmt"
	"scalable-web-services-golang-hacktiv8/module_7/gorm/database"
	"scalable-web-services-golang-hacktiv8/module_7/gorm/models"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()

	createUser("naufalaldy@gmail.com")
	// updateUserByID(1, "naufal@gmail.com")
	getUserByID(1)
	// createProduct(1, "YLO", "YOLY")
	// getUsersWithProducts()
	// deleteProductById(1)
}

// * Membuat user baru
func createUser(email string) {
	db := database.GetDB()

	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating user data:", err)
		return
	}

	fmt.Println("New user data", User)
}

// * Get user dengan Id
func getUserByID(id uint) {
	db := database.GetDB()

	user := models.User{}

	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error finding user:", err)
	}

	fmt.Printf("User Data: %+v \n", user)
}

// * Update user by id
func updateUserByID(id uint, email string) {
	db := database.GetDB()

	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error
	// * Selain method Model bisa juga dengan method Table di bawah
	// db.Table("users").Where("id = ?", 1).Updates(map[string]interface{}{"email": "bowo@gmail.com"})

	if err != nil {
		fmt.Println("Error updating user data:", err)
		return
	}
	fmt.Printf("Update user's email: %+v \n", user.Email)
}

// * Function create product
func createProduct(userID uint, brand string, name string) {
	db := database.GetDB()

	Product := models.Product{
		UserID: userID,
		Brand:  brand,
		Name:   name,
	}

	err := db.Create(&Product).Error

	if err != nil {
		fmt.Println("Error creating product data:", err.Error())
		return
	}

	fmt.Println("New product data:", Product)
}

// * Function join product dan user menggunakan eager loading, dengan method Preload
func getUsersWithProducts() {
	db := database.GetDB()

	users := models.User{}
	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting user datas with products:", err.Error())
		return
	}

	fmt.Println("User data with products")
	fmt.Printf("%+v", users)
}

// * Delete product
func deleteProductById(id uint) {
	db := database.GetDB()

	product := models.Product{}
	err := db.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		fmt.Println("Error deleting product:", err.Error())
		return
	}

	fmt.Printf("Product with id %d has been sucessfully deleted", id)
}
