package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "go_db"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

func GetConnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " +
		"password=%s dbname=%s sslmode=disable",
		 host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!" + dbname)
	return db, nil
}