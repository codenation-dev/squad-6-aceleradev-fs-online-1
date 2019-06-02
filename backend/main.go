package main

import (
	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/db"
	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	//testes rui //depois vou remover essas linhas de testes ;-)
	//services.DownloadPaymentFile(2019, 3)
	//services.ExtractRarFile("./temp/remuneracao_Marco_2019.rar", "./temp/remuneracao_Marco_2019")

	db.ConnectDataBase()
	ginRouter := gin.Default()
	ginRouter = routes.StartRouter(ginRouter)
	ginRouter.Run(":3000")

}
