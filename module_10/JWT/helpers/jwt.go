package helpers

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = "rahasia"

func GenerateToken(id uint, email string) string {
	// Jwt map claims bertipe data map[string]interface{}
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	// Melakukan enkrpsi dengan algoritma HS256 yang mengembalikan nilai pointer token
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Melakukan parsing terhadap token menjadi sebuah string panjang yang nantinya akan dikirimkan ke client
	// Secret key merupakan data yang sangat kredensial karena akan digunakan untuk mengautentikasi token yang nanti nya akan dikirimkan oleh client kepada server ketika client ingin mengirimkan sebuah request yang membutuhkan proses autentikasi token. Secret key ketika membuat token harus sama dengan secret key yang digunakan untuk mengautentikasi tokennya, jika tidak sama maka token akan dianggap tidak valid dan akan menghasilkan error.

	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	// cusotm error
	errResponse := errors.New("SIGN IN TO PROCEED")
	// Get "Authorization" header value from request
	headerToken := c.Request.Header.Get("Authorization")
	// Check if request header contains "Bearer"
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	// Get token without "Bearer" prefix
	stringToken := strings.Split(headerToken, " ")[1]

	// Parse token into jwt token pointer.  Also check if token is encrypted with HS256 algorithm
	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}
