package middleware

import (
	"gateway-golang/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Auth ...
type Auth struct{}

// AuthMiddleware ...
func (a *Auth) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization invalid",
			})
			return
		}

		_, err := utils.ValidateToken(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
}
