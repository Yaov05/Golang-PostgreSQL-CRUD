package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//Conexión a la base de datos
func getConnection() *sql.DB {
	dsn := "postgres://postgres:0519@localhost:5432/gocrud?sslmode=disable" //data source name
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
