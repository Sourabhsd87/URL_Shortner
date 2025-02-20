package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sourabhsd87/URL_Shortner/config"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		sessionKey := "session:" + token

		exists, err := config.RedisClient.Exists(config.Ctx, sessionKey).Result()
		if err != nil || exists == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired session"})
			return
		}

		c.Next()
	}
}
