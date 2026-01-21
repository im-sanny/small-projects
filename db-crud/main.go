package main

import (
	"db-crud/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Person struct {
	ID      int     `json:"id" db:"id"`
	Name    *string `json:"name" db:"name"`
	Age     *int    `json:"age" db:"age"`
	Details *string `json:"details" db:"details"`
}

func Get(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(`SELECT id, name, age, details FROM people`)
	if err != nil {
		log.Printf("failed to query people %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var people []Person

	for rows.Next() {
		var p Person
		if err := rows.Scan(&p.ID, &p.Name, &p.Age, &p.Details); err != nil {
			log.Printf("failed to scan row %v", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		people = append(people, p)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(people); err != nil {
		log.Printf("failed to encode response %v", err)
	}
}

func GetById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	c, _ := strconv.Atoi(idStr)

	var m Person
	db.DB.QueryRow(`SELECT id, name, age, details FROM people WHERE id=$1`,
		c,
	).Scan(&m.ID, &m.Name, &m.Age, &m.Details)
	json.NewEncoder(w).Encode(m)
}

func Post(w http.ResponseWriter, r *http.Request) {
	var m Person
	json.NewDecoder(r.Body).Decode(&m)

	err := db.DB.QueryRow(`
	INSERT INTO people(name, age, details)
	VALUES ($1, $2, $3)
	RETURNING id;
	`, m.Name, m.Age, m.Details).Scan(&m.ID)

	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(m)
}

func Put(w http.ResponseWriter, r *http.Request) {
	s := r.PathValue("id")
	c, _ := strconv.Atoi(s)

	var m Person
	json.NewDecoder(r.Body).Decode(&m)
	err := db.DB.QueryRow(`
	UPDATE people SET name=$1, age=$2, details=$3 WHERE id=$4
	RETURNING id, name, age, details;
	`, m.Name, m.Age, m.Details, c).Scan(&m.ID, &m.Name, &m.Age, &m.Details)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(m)
}

func Patch(w http.ResponseWriter, r *http.Request) {
	s := r.PathValue("id")
	c, _ := strconv.Atoi(s)

	var m Person
	json.NewDecoder(r.Body).Decode(&m)

	err := db.DB.QueryRow(`
	UPDATE people SET
		name = COALESCE($1, name),
		age = COALESCE($2, age),
		details = COALESCE($3, details)
	WHERE id=$4
	RETURNING id, name, age, details;
	`,
		m.Name, m.Age, m.Details, c,
	).Scan(&m.ID, &m.Name, &m.Age, &m.Details)

	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(m)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var err error
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	result, err := db.DB.Exec(`DELETE FROM people WHERE id=$1;`, id)
	if err != nil {
		log.Printf("Failed to delete person: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Failed to get rows affected: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}
}

func main() {
	db.InitDB()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /person", Get)
	mux.HandleFunc("GET /person/{id}", GetById)
	mux.HandleFunc("POST /person", Post)
	mux.HandleFunc("PUT /person/{id}", Put)
	mux.HandleFunc("PATCH /person/{id}", Patch)
	mux.HandleFunc("DELETE /person/{id}", Delete)

	fmt.Printf("server running on port :3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
