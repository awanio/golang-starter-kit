package route

import (
	"gopkg.in/gin-gonic/gin.v1"
)

func AccessControl() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := "*"

		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// c.Writer.Header().Set("Access-Control-Allow-Cookies", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With, X-Prototype-Version, x-api-key")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		// c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Origin", origin)

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
