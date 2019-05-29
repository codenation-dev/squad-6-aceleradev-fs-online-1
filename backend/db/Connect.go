package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
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

func ConnectDataBase() *sql.DB {

	//DBConnection conecao publica para acessar em outros arquivos
	var db *sql.DB
	var err error

	DBINFO := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_BANCO, DB_SSL)

	println(DBINFO)

	db, err = sql.Open(DB_SORCE, DBINFO)
	if err != nil {
		log.Fatal("Erro ao conectar no banco de dados PostgreSQL. Error:", err)
		return db
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao 'pingar' banco de dados postgresql. Error:", err)
		return nil
	}
	println("PostgreSQL.Open()")

	return db

}
