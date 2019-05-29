package db

import (
	"fmt"
	"log"

	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/models"
)

var userID int
var userEmail string
var userPassword string
var userName string
var userReceiveAlert bool

//FindAllUsers retorna todos os usuarios
func FindAllUsers() []models.User {

	var listUsers []models.User

	db := ConnectDataBase()
	defer func() {
		fmt.Println("PostgreSQL.Close()")
		db.Close()
	}()

	rows, errQuery := db.Query("select usuario.* from usuario")
	if errQuery != nil {
		log.Println("db.FindAllUsers()->Erro ao executar consulta. Error:", errQuery)
	}

	for rows.Next() {

		err := rows.Scan(&userID, &userEmail, &userPassword, &userName, &userReceiveAlert)
		if err != nil {
			log.Fatal("db.FindAllUsers()->Erro ao executar consulta. Error:", err)
		} else {
			var user = models.User{
				ID:           userID,
				Email:        userEmail,
				Password:     userPassword,
				Name:         userName,
				ReceiveAlert: userReceiveAlert,
			}
			listUsers = append(listUsers, user)
		}

	}

	return listUsers
}

//FindUserByID retona usuario pelo seu id
func FindUserByID(id int) models.User {

	var user models.User

	db := ConnectDataBase()
	defer func() {
		fmt.Println("PostgreSQL.Close()")
		db.Close()
	}()

	row := db.QueryRow("select usuario.* from usuario where usuari_id = $1", id)

	err := row.Scan(&userID, &userEmail, &userPassword, &userName, &userReceiveAlert)
	if err != nil {
		log.Println("db.FindUserByID->Erro ao executar consulta. Error:", err)
	} else {
		user = models.User{
			ID:           userID,
			Email:        userEmail,
			Password:     userPassword,
			Name:         userName,
			ReceiveAlert: userReceiveAlert,
		}
	}
	return user

}

//FindUserByEmail retona usuario pelo seu email
func FindUserByEmail(email string) models.User {

	var user models.User

	db := ConnectDataBase()
	defer func() {
		fmt.Println("PostgreSQL.Close()")
		db.Close()
	}()

	row := db.QueryRow("select usuario.* from usuario where usuari_email = $1", email)

	err := row.Scan(&userID, &userEmail, &userPassword, &userName, &userReceiveAlert)
	if err != nil {
		log.Println("db.FindUserByEmail->Erro ao executar consulta. Error:", err)
	} else {
		user = models.User{
			ID:           userID,
			Email:        userEmail,
			Password:     userPassword,
			Name:         userName,
			ReceiveAlert: userReceiveAlert,
		}
	}
	return user

}

//InsertUser retona usuario pelo seu email
func InsertUser(user models.User) models.User {

	var userUpdated models.User

	db := ConnectDataBase()
	defer func() {
		fmt.Println("PostgreSQL.Close()")
		db.Close()
	}()

	insert :=
		`INSERT INTO public.usuario
		(usuari_email, usuari_password, usuari_nome, usuari_recebe_alerta)
		VALUES ($1, $2, $3, $4) returning usuari_id, usuari_email, usuari_password, usuari_nome, usuari_recebe_alerta;`

	errUpdate := db.QueryRow(insert,
		user.Email, user.Password, user.Name, user.ReceiveAlert).Scan(&userID, &userEmail, &userPassword, &userName, &userReceiveAlert)

	if errUpdate != nil {
		log.Println("db.UpdateUserByID->Erro ao executar insert. Error:", errUpdate)
	} else {
		userUpdated = models.User{
			ID:           userID,
			Email:        userEmail,
			Password:     userPassword,
			Name:         userName,
			ReceiveAlert: userReceiveAlert,
		}
	}
	return userUpdated

}

//UpdateUserByID retona usuario pelo seu email
func UpdateUserByID(id int, user models.User) models.User {

	var userUpdated models.User

	db := ConnectDataBase()
	defer func() {
		fmt.Println("PostgreSQL.Close()")
		db.Close()
	}()

	fmt.Println(id)

	errUpdate := db.QueryRow(
		"update usuario set "+
			" email = $2 ,"+
			" password = $3 ,"+
			" nome = $4 ,"+
			" recebe_alerta = $5"+
			" where usuari_id = $1"+
			" returning usuari_id, email, password, nome, recebe_alerta;",

		id, user.Email, user.Password, user.Name, user.ReceiveAlert).Scan(&userID, &userEmail, &userPassword, &userName, &userReceiveAlert)

	if errUpdate != nil {
		log.Println("db.UpdateUserByID->Erro ao executar update. Error:", errUpdate)
	} else {
		userUpdated = models.User{
			ID:           userID,
			Email:        userEmail,
			Password:     userPassword,
			Name:         userName,
			ReceiveAlert: userReceiveAlert,
		}
	}
	return userUpdated

}

//DeleteUserByID retona usuario pelo seu email
func DeleteUserByID(id int) bool {

	db := ConnectDataBase()
	defer func() {
		fmt.Println("PostgreSQL.Close()")
		db.Close()
	}()

	fmt.Println(id)

	_, err := db.Exec("delete from usuario where usuari_id = $1", id)

	if err != nil {
		log.Println("db.DeleteUserByID->Erro ao executar delete. Error:", err)
	} else {
		return true
	}
	return false

}
