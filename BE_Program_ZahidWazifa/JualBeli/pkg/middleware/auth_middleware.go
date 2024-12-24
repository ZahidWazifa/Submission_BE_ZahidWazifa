package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware adalah middleware untuk memeriksa token JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mengambil token dari header Authorization
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak ditemukan"})
			c.Abort()
			return
		}

		// Memisahkan token dari kata "Bearer"
		tokenString := strings.Split(authHeader, "Bearer ")[1]

		// Memverifikasi token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Memastikan bahwa metode signing adalah HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("metode signing tidak valid")
			}
			// Kembalikan kunci signing
			return []byte("secret_key"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			c.Abort()
			return
		}

		// Melanjutkan ke handler berikutnya jika token valid
		c.Next()
	}
}