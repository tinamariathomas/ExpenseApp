package repo

import ("github.com/jmoiron/sqlx")

type ExpenseRepo interface{
	Insert(db *sqlx.DB,description string, amount int) (int,error)
}

const(
	InsertExpenseQuery = "INSERT INTO expense(description, amount) VALUES($1,$2) returning id"
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