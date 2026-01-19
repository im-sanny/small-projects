package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Person struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Details string `json:"details"`
}

var store = []Person{
	{ID: 1, Name: "Tom", Age: 23, Details: "Tom is a good cat"},
	{ID: 2, Name: "Jerry", Age: 22, Details: "Jerry is naughty mouse"},
	{ID: 3, Name: "Bulldog", Age: 30, Details: "Bulldog is a stupid dog"},
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewEncoder(w).Encode(store); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var p Person
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	nextID := len(store) + 1
	p.ID = nextID
	store = append(store, p)

	json.NewEncoder(w).Encode(p)
	w.WriteHeader(http.StatusCreated)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, v := range store {
		if v.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(v)
			return
		}
	}

	http.Error(w, "Person not found", http.StatusNotFound)
}

func Put(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var p Person
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i := range store {
		if store[i].ID == id {
			p.ID = id
			store[i] = p

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Person not found", http.StatusNotFound)
}

func Patch(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var patch Person
	if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i := range store {
		if store[i].ID == id {
			if patch.Name != "" {
				store[i].Name = patch.Name
			}
			if patch.Age != 0 {
				store[i].Age = patch.Age
			}
			if patch.Details != "" {
				store[i].Details = patch.Details
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(store[i])
			return
		}
	}

	http.Error(w, "Person not found", http.StatusNotFound)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i := range store {
		if store[i].ID == id {
			store = append(store[:i], store[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /person", http.HandlerFunc(Get))
	mux.Handle("POST /person", http.HandlerFunc(Post))
	mux.Handle("GET /person/{id}", http.HandlerFunc(GetById))
	mux.Handle("PUT /person/{id}", http.HandlerFunc(Put))
	mux.Handle("PATCH /person/{id}", http.HandlerFunc(Patch))
	mux.Handle("DELETE /person/{id}", http.HandlerFunc(Delete))

	fmt.Printf("server running on port :3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
