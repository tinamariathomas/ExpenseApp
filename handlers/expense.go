package handlers

import (
	"net/http"
	"git.expense-app.com/ExpenseApp/repo"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"encoding/json"
	"git.expense-app.com/ExpenseApp/models"
)



func AddExpenseHandler(expenseRepo repo.ExpenseRepo, db *sqlx.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		var expense models.Expense
		err = json.Unmarshal(body, &expense)
		if err != nil {
			panic(err)
		}
		newID, err := expenseRepo.Insert(db, expense.Description, expense.Amount)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		insertedExpense := models.Expense{Id:newID, Description:expense.Description, Amount: expense.Amount}
		response, err :=json.Marshal(insertedExpense)
		w.Write(response)
	}
}

func GetExpensesHandler(expenseRepo repo.ExpenseRepo, db *sqlx.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		expenses, err := expenseRepo.Select(db)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		response, err := json.Marshal(expenses)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(response)
	}
}
