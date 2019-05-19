package db

import (
	"log"

	"github.com/ruiblaese/projeto-codenation-banco-uati/models"
)

var id int
var email string
var password string
var name string
var receiveAlert bool

//GetUsers retorna todos os usuarios
func GetUsers() []models.User {

	var listUsers []models.User

	db := ConnectDataBase()
	defer db.Close()

	rows, errQuery := db.Query("select usuario.* from usuario")
	if errQuery != nil {
		log.Fatal("Erro ao executar consulta. Error:", errQuery)
	}

	for rows.Next() {

		err := rows.Scan(&id, &email, &password, &name, &receiveAlert)
		if err != nil {
			log.Fatal("Erro ao executar consulta. Error:", err)
		} else {
			var user = models.User{
				ID:           id,
				Email:        email,
				Password:     password,
				Name:         name,
				ReceiveAlert: receiveAlert,
			}
			listUsers = append(listUsers, user)
		}

	}

	return listUsers
}

//GetUserById retona usuario pelo seu id
func GetUserByID(id int) models.User {

	var user models.User

	db := ConnectDataBase()
	defer db.Close()

	row := db.QueryRow("select usuario.* from usuario where id = $1", id)

	err := row.Scan(&id, &email, &password, &name, &receiveAlert)
	if err != nil {
		log.Fatal("Erro ao executar consulta. Error:", err)
	} else {
		user = models.User{
			ID:           id,
			Email:        email,
			Password:     password,
			Name:         name,
			ReceiveAlert: receiveAlert,
		}
	}
	return user

}

//GetUserByEmail retona usuario pelo seu email
func GetUserByEmail(email string) models.User {

	var user models.User

	db := ConnectDataBase()
	defer db.Close()

	row := db.QueryRow("select usuario.* from usuario where email = $1", email)

	err := row.Scan(&id, &email, &password, &name, &receiveAlert)
	if err != nil {
		log.Fatal("Erro ao executar consulta. Error:", err)
	} else {
		user = models.User{
			ID:           id,
			Email:        email,
			Password:     password,
			Name:         name,
			ReceiveAlert: receiveAlert,
		}
	}
	return user

}
