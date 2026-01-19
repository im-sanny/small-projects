package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const connStr = "postgres://postgres:360420@localhost:5432/crud?sslmode=disable"

var db *sql.DB

func InitDB() {
	var err error

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("db connection successful")
}
