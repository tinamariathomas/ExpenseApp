package main

import (
	"fmt"
	"log"
	"net/http"

	"git.expense-app.com/ExpenseApp/config"
	"git.expense-app.com/ExpenseApp/router"
)

func main() {
	r := router.GetAPIRoutes()

	appConfig, err := config.ReadConfig("config.json")
	if err != nil {
		log.Fatal("Failed to read app config: ", err)
	}

	fmt.Println("Starting on port :", appConfig.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", appConfig.Port), r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}


