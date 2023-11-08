package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		if name != "" {
			c.AbortWithStatusJSON(200, gin.H{
				"status": 401,
			})
		}
		c.Next()
	}
}
