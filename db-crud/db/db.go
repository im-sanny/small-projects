package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const connStr = "postgres://postgres:360420@localhost:5432/crud?sslmode=disable"

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("db connection successful")
}
