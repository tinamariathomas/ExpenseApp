package dbtest

import (
	"testing"

	"git.expense-app.com/ExpenseApp/repo"
	"github.com/stretchr/testify/assert"
)

func TestInsertExpenseSuccess(t *testing.T){
	db := initTestDB()
	defer db.Close()


	expense := repo.Expense{}
	id, err := expense.Insert(db, "Books", 123)

	assert.NotEqual(t, 0, id)
	assert.NoError(t,err)
}

func TestSelectExpensesSuccess(t *testing.T){
	db := initTestDB()
	defer db.Close()

	expense := repo.Expense{}
	expense.Insert(db, "Books", 123)
	expense.Insert(db, "Arial soap", 500)

	expenses, err := expense.Select(db)

	assert.NoError(t,err)
	assert.Equal(t, 2, len(expenses))
}
