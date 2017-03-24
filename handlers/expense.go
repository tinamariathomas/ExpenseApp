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
		expenseRepo.Insert(db, expense.Description, expense.Amount)
		w.Write([]byte("Server check. I'm alive!"))
	}
}
