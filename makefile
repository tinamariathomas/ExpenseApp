.PHONY:build
build:
	 go build -o out/build/ExpenseApp

.PHONY:db_create
db_create:
	psql -c "CREATE DATABASE expense_manager"

.PHONY:db_migrate
db_migrate:
	goose up
