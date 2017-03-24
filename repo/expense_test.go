package repo_test

import (
	"testing"

	"git.expense-app.com/ExpenseApp/repo"
	"github.com/stretchr/testify/assert"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"errors"
	"regexp"
	"fmt"
)

func TestInsertExpenseSuccess(t *testing.T){

	db, mock, err := sqlmock.New()
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(1521)
	mock.ExpectQuery(removeLines(repo.InsertExpenseQuery)).WithArgs("adidas shoes",10).WillReturnRows(rows)

	expense := repo.Expense{}

	id, err := expense.Insert(sqlx.NewDb(db, "postgres"), "adidas shoes", 10)

	assert.NoError(t, err)
	assert.Equal(t, 1521, id)
}

func TestInsertExpenseFailsForDBError(t *testing.T){

	db, mock, err := sqlmock.New()
	mock.ExpectQuery(repo.InsertExpenseQuery).WithArgs("adidas shoes", 10).WillReturnError(errors.New("Much to learn, you still have"))

	expense := repo.Expense{}

	id, err := expense.Insert(sqlx.NewDb(db, "postgres"), "adidas shoes", 10)

	assert.Error(t, err)
	assert.Equal(t, 0, id)
}

func TestSelectExpenseSuccess(t *testing.T){

	db, mock, err := sqlmock.New()
	rows := sqlmock.NewRows([]string{"id","description", "amount"})
	rows.AddRow(1521, "shoes - adidas", 4500)
	rows.AddRow(1324, "oreo milkshake", 35)

	mock.ExpectQuery(removeLines(repo.SelectExpensesQuery)).WillReturnRows(rows)

	expense := repo.Expense{}

	expenses, err := expense.Select(sqlx.NewDb(db, "postgres"))

	assert.NoError(t, err)
	assert.Equal(t, 2, len(expenses))
}

func TestSelectExpenseFailsForDBError(t *testing.T){

	db, mock, err := sqlmock.New()

	mock.ExpectQuery(removeLines(repo.SelectExpensesQuery)).WillReturnError(errors.New("For every job that must be done, there is an element of fun"))

	expense := repo.Expense{}

	expenses, err := expense.Select(sqlx.NewDb(db, "postgres"))

	assert.Error(t, err)
	assert.Nil(t, expenses)
}


func removeLines(query string) string {
	r, err := regexp.Compile("[\n]+")
	if err != nil {
		fmt.Println("Something went wrong")
	}
	query = r.ReplaceAllString(query, " ")
	return regexp.QuoteMeta(query)
}