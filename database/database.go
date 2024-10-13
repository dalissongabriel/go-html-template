package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnDatabase() *sql.DB {
	connStr := "user=root password=root dbname=root port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}

	return db
}
