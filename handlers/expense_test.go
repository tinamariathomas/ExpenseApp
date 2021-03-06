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
	"errors"
	"io/ioutil"
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
	
	insertedExpense := &models.Expense{}
	responseBody, err := ioutil.ReadAll(w.Body)
	require.NoError(t, err)

	json.Unmarshal(responseBody, insertedExpense)

	assert.Equal(t,http.StatusOK, w.Code)
	assert.Equal(t, 123, insertedExpense.Id)
	mockExpense.AssertExpectations(t)
}

func TestInsertExpenseHandlerFailsForDBInsert(t *testing.T) {
	expense := models.Expense{Description:"shoes", Amount:1000}
	mockExpense := &MockExpense{}
	mockDB := &sqlx.DB{}
	mockExpense.On("Insert", mockDB, "shoes",1000).Return(0, errors.New("No Luke. I am your Father"))

	requestBody ,err := json.Marshal(expense)
	require.NoError(t, err)

	r, err := http.NewRequest("POST", "", bytes.NewBuffer(requestBody))
	require.NoError(t, err)
	w := httptest.NewRecorder()

	h.AddExpenseHandler(mockExpense, mockDB)(w,r)

	assert.Equal(t,http.StatusInternalServerError, w.Code)
	mockExpense.AssertExpectations(t)
}

func TestSelectExpensesHandlerSuccess(t *testing.T) {
	expenses := []models.Expense{{Id: 12, Description:"shoes", Amount:1000}}

	mockExpense := &MockExpense{}
	mockDB := &sqlx.DB{}
	mockExpense.On("Select", mockDB).Return(expenses, nil)

	r, err := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()

	h.GetExpensesHandler(mockExpense, mockDB)(w,r)

	responseBody, err := ioutil.ReadAll(w.Body)
	require.NoError(t, err)

	responseExpenses := []models.Expense{}
	json.Unmarshal(responseBody, &responseExpenses)

	assert.Equal(t,http.StatusOK, w.Code)
	assert.Equal(t, 1, len(responseExpenses))
	assert.Equal(t, expenses, responseExpenses)
	mockExpense.AssertExpectations(t)
}

func TestSelectExpensesHandlerFailsForDBError(t *testing.T) {
	mockExpense := &MockExpense{}
	mockDB := &sqlx.DB{}
	mockExpense.On("Select", mockDB).Return(nil, errors.New("Frankly my dear, I don't give a damn"))

	r, err := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()

	h.GetExpensesHandler(mockExpense, mockDB)(w,r)

	responseBody, err := ioutil.ReadAll(w.Body)
	require.NoError(t, err)

	responseExpenses := []models.Expense{}
	json.Unmarshal(responseBody, &responseExpenses)

	assert.Equal(t,http.StatusInternalServerError, w.Code)
	assert.Equal(t, 0, len(responseExpenses))

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


func (m *MockExpense) Select(db *sqlx.DB) (expenses []models.Expense,err error){
	args := m.Called(db)

	if args[0] != nil {
		expenses = args[0].([]models.Expense)
	}

	if args[1] != nil {
		err = args[1].(error)
	}

	return expenses, err
}