package handlers

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/db"
	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/models"
	"github.com/gin-gonic/gin"
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
		c.AbortWithStatus(http.StatusNoContent)
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
		c.JSON(http.StatusNoContent, gin.H{"code": "ERROR", "message": "Invalid param"})
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

	fileTemp := "./temp/" + file.Filename
	c.SaveUploadedFile(file, fileTemp)

	go registerCustomersFromCSV(fileTemp)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func registerCustomersFromCSV(file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Panicln(err)
	}
	defer f.Close() // this needs to be after the err check

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Panicln(err)
	}

	for _, line := range lines {
		nomeCliente := line[0]
		fmt.Println(nomeCliente)

		custumer := db.FindCustomerByName(nil, nomeCliente)

		if custumer.ID > 0 {
			//talvez fazer update com data ultima atualizacao, talvez nome ultimo arquivo
			//db.UpdateCustomerByID(custumer.ID, models.Customer{Name: nomeCliente})
		} else {
			db.InsertCustomer(models.Customer{Name: nomeCliente})
		}

	}
}
