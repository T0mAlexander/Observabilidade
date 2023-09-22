package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DatabaseConnection() *sql.DB {
	connection := "user=golang dbname=postgresql password=golang host=database sslmode=disable"

	database, error := sql.Open("postgres", connection)

	if error != nil {
		panic(error.Error())
	}

	return database
}
