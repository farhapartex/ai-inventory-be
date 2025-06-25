package middlewares

import (
	"net/http"
	"strings"

	"github.com/farhapartex/ainventory/config"
	"github.com/farhapartex/ainventory/models"
	"github.com/farhapartex/ainventory/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}

		var user models.User
		result := config.DB.Where("id = ?", claims.UserID).First(&user)
		if result.Error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		if user.Status != "active" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Account is not active"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Set("userId", user.ID)

		c.Next()
	}
}
