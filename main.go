package main

import (
	"crmservice/modules/account"
	"crmservice/modules/customer"
	"crmservice/utils/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	dbCrud := db.GormMysql()
	fmt.Println("Database connected successfully")
	customerHandler := customer.NewRouter(dbCrud)
	customerHandler.Handle(router)

	accountHandler := account.NewRouter(dbCrud)
	accountHandler.Handle(router)

	port := ":8082"
	errRoute := router.Run(port)
	if errRoute != nil {
		fmt.Println("Error running server", errRoute)
		return
	}
}
