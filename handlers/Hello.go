package handlers

import (
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/ruiblaese/projeto-codenation-banco-uati/db"
)

//HelloHandler teste de login
func HelloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	var user = db.FindUserByEmail(claims["email"].(string))
	c.JSON(200, gin.H{
		"email": claims["email"],
		"user":  user,
		"text":  "Hello World.",
	})
}
