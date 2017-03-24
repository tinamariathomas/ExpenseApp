package repo

import ("github.com/jmoiron/sqlx"
	"git.expense-app.com/ExpenseApp/models"
)

type ExpenseRepo interface{
	Insert(db *sqlx.DB,description string, amount int) (int,error)
	Select(db *sqlx.DB) ([]models.Expense, error)
}

const(
	InsertExpenseQuery = "INSERT INTO expense(description, amount) VALUES($1,$2) returning id"
	SelectExpensesQuery = "SELECT * from expense"
)

type Expense struct{

}

func (e *Expense) Insert(db *sqlx.DB,description string, amount int) (int, error){

	var id int
	rows, err := db.Query(InsertExpenseQuery, description, amount)
	if err != nil{
		return 0, err
	}

	if rows.Next() {
		rows.Scan(&id)
	}

	return int(id), err
}

func (e *Expense) Select(db *sqlx.DB) ([]models.Expense, error){
	var expenses []models.Expense
	err := db.Select(&expenses, SelectExpensesQuery)
	if err != nil{
		return nil, err
	}
	return expenses, nil
}