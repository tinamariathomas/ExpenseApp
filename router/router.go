package router

import (
	"github.com/gorilla/mux"

	h "git.expense-app.com/ExpenseApp/handlers"
)

func GetAPIRoutes() *mux.Router{

	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", h.HealthCheckHandler).Methods("GET")

	return r
}
