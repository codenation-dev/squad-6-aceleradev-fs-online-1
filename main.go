package main

import (
	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/db"
	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/routes"
)

func main() {

	db.ConnectDataBase()

	routes.StartRouter()

}
