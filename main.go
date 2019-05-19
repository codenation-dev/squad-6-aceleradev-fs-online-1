package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/ruiblaese/projeto-codenation-banco-uati/routes"
)

var (
	DB_HOST     = os.Getenv("DB_HOST")
	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_BANCO    = os.Getenv("DB_BANCO")
	DB_PORT     = os.Getenv("DB_PORT")

	DB_SSL   = "disable"
	DB_SORCE = "postgres"
)

var (
	db  *sql.DB
	err error
)

func init() {

	DBINFO := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_BANCO, DB_SSL)

	println(DBINFO)

	db, err = sql.Open(DB_SORCE, DBINFO)
	if err != nil {
		// log.Println(err)
		return
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		//log.Println(err)
		return
	}
	println("Banco de dados PostgreSQL iniciando com sucesso!")
	return
}

func main() {

	routes.StartRouter()

}
