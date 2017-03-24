package router

import (
	h "git.expense-app.com/ExpenseApp/handlers"
	repo "git.expense-app.com/ExpenseApp/repo"

	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"fmt"
)

func GetAPIRoutes() *mux.Router{

	r := mux.NewRouter()
	expense := &repo.Expense{}

	db := openDBConnection("<username>","<db-name>","<password>")

	insertExpenseHandler := h.AddExpenseHandler(expense, db)
	getExpensesHandler := h.GetExpensesHandler(expense,db)

	r.HandleFunc("/healthcheck", h.HealthCheckHandler).Methods("GET")
	r.HandleFunc("/expense", insertExpenseHandler).Methods("POST")
	r.HandleFunc("/expense", getExpensesHandler).Methods("GET")

	return r
}

func openDBConnection(username string, dbName string, password string) *sqlx.DB{
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", username, dbName, password))
	if err != nil{
		panic(err)
	}
	return db
}
