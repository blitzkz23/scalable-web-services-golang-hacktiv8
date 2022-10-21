package middlewares

import (
	"net/http"
	"scalable-web-services-golang-hacktiv8/module_10/JWT/helpers"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Panggil fungsi VerifyToken untuk melakukan verifikasi token
		verifyToken, err := helpers.VerifyToken(c)

		if err != nil {
			// Untuk membatalkan middleware perlu abort
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}
		// Menyimpan claim dari token ke dalam data request agar bisa diambil pada endpoint
		c.Set("userData", verifyToken)
		c.Next()
	}
}
