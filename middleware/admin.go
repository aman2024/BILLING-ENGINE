package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		authToken := c.Request.Header.Get("Authorization")

		if authToken == "" {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"status": false, "message": "Empty Authorization Token",
			})
			return
		}

		c.Set("adminId", authToken)
		c.Next()
	}
}
