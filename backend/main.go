package main

import (
	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/db"
	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db.ConnectDataBase()

	ginRouter := gin.Default()
	ginRouter = routes.StartRouter(ginRouter)
	ginRouter.Run(":4000")
}
