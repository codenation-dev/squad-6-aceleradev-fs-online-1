package middleware

import "github.com/gin-gonic/gin"

//CorsMiddleware middleware para liberar cors
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//if c.Request.Method != "POST" {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		//c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		//c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		//c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		//c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
		/*
			} else {
				c.Next()
			}
		*/

	}
}
