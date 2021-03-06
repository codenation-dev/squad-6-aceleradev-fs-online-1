package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/models"
)

//FindAllCustomers retorna todos os usuarios
func FindAllCustomers() []models.Customer {
	var (
		customerID    int
		customerName  string
		listCustomers []models.Customer
	)

	db := ConnectDataBase()
	defer CloseDataBase(db)

	rows, errQuery := db.Query("select cliente.* from cliente order by client_id")
	if errQuery != nil {
		log.Println("db.FindAllCustomers()->Erro ao executar consulta. Error:", errQuery)
	}

	for rows.Next() {
		err := rows.Scan(&customerID, &customerName)
		if err != nil {
			log.Fatal("db.FindAllCustomers()->Erro ao executar consulta. Error:", err)
		} else {
			var customer = models.Customer{
				ID:   customerID,
				Name: customerName}
			listCustomers = append(listCustomers, customer)
		}

	}

	return listCustomers
}

//FindCustomerByID retona usuario pelo seu id
func FindCustomerByID(id int) models.Customer {
	var (
		customerID   int
		customerName string
		customer     models.Customer
	)
	if id > 0 {
		db := ConnectDataBase()
		defer CloseDataBase(db)

		row := db.QueryRow("select cliente.* from cliente where client_id = $1", id)

		err := row.Scan(&customerID, &customerName)
		if (err != nil) && (err != sql.ErrNoRows) {
			log.Println("db.FindCustomerByID->Erro ao executar consulta. Error:", err)
		} else {
			customer = models.Customer{
				ID:   customerID,
				Name: customerName}
		}
	}
	return customer
}

//FindCustomerByName retona usuaclienterio pelo seu email
func FindCustomerByName(optionalDB *sql.DB, name string) models.Customer {
	var (
		customerID   int
		customerName string
		customer     models.Customer
	)

	var db *sql.DB
	if optionalDB != nil {
		db = optionalDB
	} else {
		db = ConnectDataBase()
		defer CloseDataBase(db)
	}

	row := db.QueryRow("select cliente.* from cliente where client_nome = $1", name)

	err := row.Scan(&customerID, &customerName)
	if (err != nil) && (err != sql.ErrNoRows) {
		log.Println("db.FindCustomerByName->Erro ao executar consulta. Error:", err)
	} else {
		customer = models.Customer{
			ID:   customerID,
			Name: customerName,
		}
	}
	return customer
}

//InsertCustomer retona usuario pelo seu email
func InsertCustomer(customer models.Customer) models.Customer {
	var (
		customerID       int
		customerName     string
		customerInserted models.Customer
	)

	db := ConnectDataBase()
	defer CloseDataBase(db)

	insert :=
		`INSERT INTO public.cliente
		(client_nome)
		VALUES ($1) returning client_id, client_nome;`

	errUpdate := db.QueryRow(insert,
		customer.Name).Scan(&customerID, &customerName)

	if (errUpdate != nil) && (errUpdate != sql.ErrNoRows) {
		log.Println("db.UpdateCustomerByID->Erro ao executar insert. Error:", errUpdate)
	} else {
		customerInserted = models.Customer{
			ID:   customerID,
			Name: customerName}
	}
	return customerInserted
}

//UpdateCustomerByID retona usuario pelo seu email
func UpdateCustomerByID(id int, customer models.Customer) models.Customer {
	var (
		customerID      int
		customerName    string
		customerUpdated models.Customer
	)

	db := ConnectDataBase()
	defer CloseDataBase(db)

	fmt.Println("update cliente set " +
		" client_nome = $2 ," +

		" where client_id = $1" +
		" returning client_id, client_nome;")

	errUpdate := db.QueryRow("update cliente set "+
		" client_nome = $2 "+
		" where client_id = $1"+
		" returning client_id, client_nome;",

		id, customer.Name).Scan(&customerID, &customerName)

	if (errUpdate != nil) && (errUpdate != sql.ErrNoRows) {
		log.Println("db.UpdateCustomerByID->Erro ao executar update. Error:", errUpdate)
	} else {
		customerUpdated = models.Customer{
			ID:   customerID,
			Name: customerName}
	}
	return customerUpdated
}

//DeleteCustomerByID retona usuario pelo seu email
func DeleteCustomerByID(id int) bool {
	db := ConnectDataBase()
	defer CloseDataBase(db)

	_, err := db.Exec("delete from cliente where client_id = $1", id)

	if err != nil {
		log.Println("db.DeleteCustomerByID->Erro ao executar delete. Error:", err)
	} else {
		return true
	}
	return false
}

//CountCustomers retona quantidade de usuarios
func CountCustomers() int64 {
	db := ConnectDataBase()
	defer CloseDataBase(db)

	var count int64

	row := db.QueryRow("select count(cliente.client_id) from cliente")

	err := row.Scan(&count)
	if (err != nil) && (err != sql.ErrNoRows) {
		log.Println("db.CountCustomers->Erro ao executar consulta. Error:", err)
	}
	return count
}
