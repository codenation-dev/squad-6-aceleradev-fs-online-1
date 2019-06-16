package main

import (
	"fmt"
	"time"

	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/db"
	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/handlers"
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
	fmt.Println(time.Now())
	db.DeleteAllPayment()
	handlers.CheckPayments()
}
