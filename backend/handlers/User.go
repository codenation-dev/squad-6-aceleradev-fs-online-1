package handlers

import (
	"net/http"
	"strconv"

	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/db"
	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/models"
	"github.com/gin-gonic/gin"
)

// GetUsers retorna todos os usuarios
func GetUsers(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, db.FindAllUsers())
}

// GetUser retornoa um usuario, busca primeiro por id, se nao conseguir converter para inteiro busca por email
func GetUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var user models.User

	if id, err := strconv.Atoi(c.Params.ByName("id")); err == nil {
		user = db.FindUserByID(id)

	} else {
		user = db.FindUserByEmail(c.Params.ByName("id"))
	}

	if user.ID > 0 {
		c.JSON(http.StatusOK, user)
	} else {
		c.AbortWithStatus(http.StatusNoContent)
	}
}

// PutUser atualiza informacoes do usuario
func PutUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user models.User
	c.Bind(&user)

	if id, err := strconv.Atoi(c.Params.ByName("id")); err == nil {
		userUpdated := db.UpdateUserByID(id, user)

		if userUpdated.ID > 0 {
			c.JSON(http.StatusCreated, userUpdated)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "ERROR", "message": "Internal Server Error"})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": "ERROR", "message": "Invalid param"})
	}
}

// NewUser cria novo usuario
func NewUser(c *gin.Context) {
	//c.Header("Content-Type", "application/json")

	var user models.User
	c.Bind(&user)

	userInserted := db.InsertUser(user)

	if userInserted.ID > 0 {
		c.JSON(http.StatusOK, userInserted)

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": "ERROR", "message": "Internal Server Error"})
	}
}

// DeleteUser deleta usuario
func DeleteUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	if id, err := strconv.Atoi(c.Params.ByName("id")); err == nil {
		apagou := db.DeleteUserByID(id)

		if apagou {
			c.JSON(http.StatusOK, gin.H{"code": "OK", "message": ""})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "ERROR", "message": "Internal Server Error"})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": "ERROR", "message": "Invalid param"})
	}
}
