package main

import (
	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/db"
	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	//handlers.CheckPayments()

	db.ConnectDataBase()
	ginRouter := gin.Default()
	ginRouter = routes.StartRouter(ginRouter)
	ginRouter.Run(":3000")

}
