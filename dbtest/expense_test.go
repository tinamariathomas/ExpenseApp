package dbtest

import (
	"testing"

	"git.expense-app.com/ExpenseApp/repo"
	"github.com/stretchr/testify/assert"
)

func TestInsertExpenseSuccess(t *testing.T){
	db := initTestDB()


	expense := repo.Expense{}
	id, err := expense.Insert(db, "Books", 123)

	assert.NotEqual(t, 0, id)
	assert.NoError(t,err)
}
