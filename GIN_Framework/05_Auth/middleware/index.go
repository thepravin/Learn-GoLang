package middleware

import (
	"ginLearning/05_Auth/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckMiddleware(c *gin.Context) {
	headers := c.GetHeader("Authorization")

	if headers == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization header required",
		})
		return
	}

	// Split "Bearer <token>" format
	tokenParts := strings.Split(headers, " ")

	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid authorization format. Use: Bearer <token>",
		})
		return
	}

	// Validate token and check expiration
	claims, err := utils.TokenCheck(tokenParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Store claims in context for use in handlers
	c.Set("user_email", claims["email"])
	c.Set("user_id", claims["id"])

	c.Next()
}
