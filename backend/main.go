package main

import (
	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/db"
	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/routes"
)

func main() {

	//testes rui //depois vou remover ;-)
	//services.DownloadPaymentFile(2019, 3)

	db.ConnectDataBase()

	routes.StartRouter()

}
