package middlewares

import (
	"github.com/gin-gonic/gin"
)

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}

func HasPermissions() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
