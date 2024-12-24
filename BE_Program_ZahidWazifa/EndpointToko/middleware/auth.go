package middleware

import (
	"strings"
	"task3/utils"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware mengecek validitas JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if (authHeader == "") {
			c.JSON(401, gin.H{"error": "Authorization header tidak ditemukan"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "Format token tidak valid"})
			c.Abort()
			return
		}

		userID, err := utils.ValidateToken(parts[1])
		if err != nil {
			c.JSON(401, gin.H{"error": "Token tidak valid"})
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
