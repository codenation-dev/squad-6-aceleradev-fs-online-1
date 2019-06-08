package main

import (
	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/services"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	/*
		db.ConnectDataBase()

		go handlers.MonitorPayments()

		ginRouter := gin.Default()
		ginRouter = routes.StartRouter(ginRouter)

		ginRouter.Run(":4000")
	*/
	services.SendEmailNotificationNewsEmployee()
}
