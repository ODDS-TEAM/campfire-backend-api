init:
	sqlite3 campfire.db < _scripts/init.sql
run:
	go run main.go
test:
	go test -v ./...
