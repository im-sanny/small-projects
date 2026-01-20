package main

import (
	"db-crud/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	ID      int     `json:"id" db:"id"`
	Name    *string `json:"name" db:"name"`
	Age     *int    `json:"age" db:"age"`
	Details *string `json:"details" db:"details"`
}

func Get(w http.ResponseWriter, r *http.Request) {
	rows, _ := db.DB.Query(`SELECT id, name, age, details FROM people`)
	defer rows.Close()

	var people []Person

	for rows.Next() {
		var p Person
		rows.Scan(&p.ID, &p.Name, &p.Age, &p.Details)
		people = append(people, p)
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	db.InitDB()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /person", Get)

	fmt.Printf("server running on port :3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
