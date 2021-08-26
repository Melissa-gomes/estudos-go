package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectWithDataBase() *sql.DB {
	connection := "user=postgres dbname=alura_loja password=codenation host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}
