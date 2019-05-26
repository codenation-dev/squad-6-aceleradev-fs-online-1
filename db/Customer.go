package db

import (
	"fmt"
	"log"

	"github.com/ruiblaese/projeto-codenation-banco-uati/models"
)

var customerID int
var customerName string

//FindAllCustomers retorna todos os usuarios
func FindAllCustomers() []models.Customer {

	var listCustomers []models.Customer

	db := ConnectDataBase()
	defer func() {
		fmt.Println("fechou conexao com postgresql")
		db.Close()
	}()

	rows, errQuery := db.Query("select cliente.* from cliente")
	if errQuery != nil {
		log.Println("db.FindAllCustomers()->Erro ao executar consulta. Error:", errQuery)
	}

	for rows.Next() {

		err := rows.Scan(&customerID, &customerName)
		if err != nil {
			log.Fatal("db.FindAllCustomers()->Erro ao executar consulta. Error:", err)
		} else {
			var Customer = models.Customer{
				ID:   customerID,
				Name: customerName}
			listCustomers = append(listCustomers, Customer)
		}

	}

	return listCustomers
}

//FindCustomerByID retona usuario pelo seu id
func FindCustomerByID(id int) models.Customer {

	var Customer models.Customer

	db := ConnectDataBase()
	defer func() {
		fmt.Println("fechou conexao com postgresql")
		db.Close()
	}()

	row := db.QueryRow("select cliente.* from cliente where usuari_id = $1", id)

	err := row.Scan(&customerID, &customerName)
	if err != nil {
		log.Println("db.FindCustomerByID->Erro ao executar consulta. Error:", err)
	} else {
		Customer = models.Customer{
			ID:   customerID,
			Name: customerName}
	}
	return Customer

}

//InsertCustomer retona usuario pelo seu email
func InsertCustomer(Customer models.Customer) models.Customer {

	var CustomerUpdated models.Customer

	db := ConnectDataBase()
	defer func() {
		fmt.Println("fechou conexao com postgresql")
		db.Close()
	}()

	insert :=
		`INSERT INTO public.cliente
		(client_nome)
		VALUES ($1) returning client_id, client_nome;`

	errUpdate := db.QueryRow(insert,
		Customer.Name).Scan(&customerID, &customerName)

	if errUpdate != nil {
		log.Println("db.UpdateCustomerByID->Erro ao executar insert. Error:", errUpdate)
	} else {
		CustomerUpdated = models.Customer{
			ID:   customerID,
			Name: customerName}
	}
	return CustomerUpdated

}

//UpdateCustomerByID retona usuario pelo seu email
func UpdateCustomerByID(id int, Customer models.Customer) models.Customer {

	var CustomerUpdated models.Customer

	db := ConnectDataBase()
	defer func() {
		fmt.Println("fechou conexao com postgresql")
		db.Close()
	}()

	fmt.Println(id)

	errUpdate := db.QueryRow(
		"update cliente set "+
			" client_nome = $2 ,"+

			" where client_id = $1"+
			" returning client_id, client_nome;",

		id, Customer.Name).Scan(&customerID, &customerName)

	if errUpdate != nil {
		log.Println("db.UpdateCustomerByID->Erro ao executar update. Error:", errUpdate)
	} else {
		CustomerUpdated = models.Customer{
			ID:   customerID,
			Name: customerName}
	}
	return CustomerUpdated

}

//DeleteCustomerByID retona usuario pelo seu email
func DeleteCustomerByID(id int) bool {

	db := ConnectDataBase()
	defer func() {
		fmt.Println("fechou conexao com postgresql")
		db.Close()
	}()

	fmt.Println(id)

	_, err := db.Exec("delete from cliente where client_id = $1", id)

	if err != nil {
		log.Println("db.DeleteCustomerByID->Erro ao executar delete. Error:", err)
	} else {
		return true
	}
	return false

}
