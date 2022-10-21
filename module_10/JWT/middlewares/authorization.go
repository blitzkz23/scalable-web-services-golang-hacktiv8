package middlewares

import (
	"net/http"
	"scalable-web-services-golang-hacktiv8/module_10/JWT/database"
	"scalable-web-services-golang-hacktiv8/module_10/JWT/models"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		// Memastikan ada parameter id di url dan valid
		productID, err := strconv.Atoi(c.Param("productID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid parameter",
			})
			return
		}
		// Mendapatkan token yang disimpan auth
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Product := models.Product{}

		// Cek apakah product dengan id tersebut ada dan milik user yang sedang login
		err = db.Select("user_id").First(&Product, uint(productID)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "Product not found",
			})
			return
		}

		if Product.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You don't have permission to access this product",
			})
			return
		}

		c.Next()
	}
}
