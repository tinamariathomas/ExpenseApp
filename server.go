package main

import (
	"fmt"
	"log"
	"net/http"

	"git.expense-app.com/ExpenseApp/config"
)

func main() {
	http.HandleFunc("/", helloWorld)

	appConfig, err := config.ReadConfig("config.json")
	if err != nil {
		log.Fatal("Failed to read app config: ", err)
	}

	fmt.Println("Starting on port :", appConfig.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", appConfig.Port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
