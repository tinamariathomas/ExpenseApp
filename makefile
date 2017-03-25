.PHONY:build
build:
	 go build -o out/build/ExpenseApp

GoFolders = $(shell find . -name '*.go' -not -path "./vendor/*" | grep -o '\/[a-z]\+\/' | grep -o '[a-z]\+' | uniq)
.PHONY:test
test:
	for number in $(GoFolders) ; do \
        	go test './'$$number ; \
	done

.PHONY:db_create
db_create:
	psql -c "CREATE DATABASE expense_manager"

.PHONY:db_migrate
db_migrate:
	goose up
