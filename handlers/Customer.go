package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/db"
	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/models"
)

// GetCustomers retorna todos os usuarios
func GetCustomers(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, db.FindAllCustomers())
}

// GetCustomer retornoa um usuario, busca primeiro por id, se nao conseguir converter para inteiro busca por email
func GetCustomer(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var customer models.Customer

	if id, err := strconv.Atoi(c.Params.ByName("id")); err == nil {
		customer = db.FindCustomerByID(id)
	}

	if customer.ID > 0 {
		c.JSON(http.StatusOK, customer)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}

}

// PutCustomer atualiza informacoes do usuario
func PutCustomer(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var customer models.Customer
	c.Bind(&customer)

	if id, err := strconv.Atoi(c.Params.ByName("id")); err == nil {

		CustomerUpdated := db.UpdateCustomerByID(id, customer)

		if CustomerUpdated.ID > 0 {
			c.JSON(http.StatusOK, CustomerUpdated)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "ERROR", "message": "Internal Server Error"})
		}

	} else {
		c.JSON(http.StatusNotFound, gin.H{"code": "ERROR", "message": "Invalid param"})
	}

}

// NewCustomer cria novo usuario
func NewCustomer(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var customer models.Customer
	c.Bind(&customer)

	CustomerInserted := db.InsertCustomer(customer)

	if CustomerInserted.ID > 0 {
		c.JSON(http.StatusOK, CustomerInserted)

	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "ERROR", "message": "Internal Server Error"})
	}

}

// DeleteCustomer deleta usuario
func DeleteCustomer(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	if id, err := strconv.Atoi(c.Params.ByName("id")); err == nil {

		apagou := db.DeleteCustomerByID(id)

		if apagou {
			c.JSON(http.StatusOK, gin.H{"code": "OK", "message": ""})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "ERROR", "message": "Internal Server Error"})
		}

	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "ERROR", "message": "Internal Server Error"})
	}

}

// UploadCustomersWithCSV atualiza informacoes do usuario
func UploadCustomersWithCSV(c *gin.Context) {

	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// Upload the file to specific dst.
	c.SaveUploadedFile(file, "./temp/"+file.Filename)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

}
