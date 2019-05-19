package main

import (
	"github.com/ruiblaese/projeto-codenation-banco-uati/db"
	"github.com/ruiblaese/projeto-codenation-banco-uati/routes"
)

func main() {

	db.ConnectDataBase()

	routes.StartRouter()

}
