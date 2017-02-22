package dbtest

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	_ "github.com/lib/pq"

)

func initTestDB() *sqlx.DB{
	db, err := sqlx.Connect("postgres", "user=***** dbname=expense_manager sslmode=disable password=*****")
	if err != nil {
		fmt.Println("Encountered error while trying to connect to DB :",err)
		return nil
	}

	_, err = db.Exec("DROP DATABASE IF EXISTS expense_manager_test")
	if err != nil {
		fmt.Println("Encountered error while deleting test database :",err)
		return nil
	}

	_, err = db.Exec("CREATE DATABASE expense_manager_test TEMPLATE expense_manager")
	if err != nil{
		fmt.Println("Encountered error while creating test database :",err)
		return nil
	}


	return db
}