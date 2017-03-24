package handlers_test

import (
	"testing"

	h "git.expense-app.com/ExpenseApp/handlers"
	"net/http/httptest"
	"net/http"
	"github.com/stretchr/testify/assert"
	"git.expense-app.com/ExpenseApp/models"
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/jmoiron/sqlx"
	"bytes"
)

func TestInsertExpenseHandlerSuccess(t *testing.T) {
	expense := models.Expense{Description:"shoes", Amount:1000}
	mockExpense := &MockExpense{}
	mockDB := &sqlx.DB{}
	mockExpense.On("Insert", mockDB, "shoes",1000).Return(123, nil)

	requestBody ,err := json.Marshal(expense)
	require.NoError(t, err)

	r, err := http.NewRequest("POST", "", bytes.NewBuffer(requestBody))
	require.NoError(t, err)
	w := httptest.NewRecorder()

	h.AddExpenseHandler(mockExpense, mockDB)(w,r)

	assert.Equal(t,http.StatusOK, w.Code)
	mockExpense.AssertExpectations(t)
}


type MockExpense struct {
	mock.Mock
}

func (m *MockExpense) Insert(db *sqlx.DB, description string, amount int) (int,error){
	args := m.Called(db, description, amount)
	var id int
	var err error

	if args[0] != nil {
		id = args[0].(int)
	}

	if args[1] != nil {
		err = args[1].(error)
	}

	return id, err
}
