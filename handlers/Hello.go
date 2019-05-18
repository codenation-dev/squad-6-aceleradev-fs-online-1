package handlers

import (
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/ruiblaese/projeto-codenation-banco-uati/models"
)

var identityKey = "id"

func HelloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims["id"],
		"userName": user.(*models.User).UserName,
		"text":     "Hello World.",
	})
}
