package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Person struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Details string `json:"details"`
}

var Human = []Person{
	{Id: 1, Name: "Tom", Age: 23, Details: "Tom is a good cat"},
	{Id: 2, Name: "Jerry", Age: 22, Details: "Jerry is naughty mouse"},
	{Id: 3, Name: "Bulldog", Age: 30, Details: "Bulldog is a stupid dog"},
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	encoder := json.NewEncoder(w).Encode(Human)
	if encoder == nil {
		http.Error(w, "human not found", http.StatusNotFound)
		return
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	var m Person
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	nI := len(Human) + 1
	m.Id = nI
	Human = append(Human, m)
	err1 := json.NewEncoder(w).Encode(m)
	if err1 != nil {
		http.Error(w, "failed to create data", http.StatusBadRequest)
		return
	}
}

func GetById(w http.ResponseWriter, r *http.Request) {
	s := r.PathValue("id")
	c, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	for _, v := range Human {
		if v.Id == c {
			err := json.NewEncoder(w).Encode(v)
			if err != nil {
				http.Error(w, "data not found", http.StatusNotFound)
				return
			}
		}
	}
	w.WriteHeader(http.StatusAccepted)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /person", http.HandlerFunc(Get))
	mux.Handle("POST /person", http.HandlerFunc(Post))
	mux.Handle("GET /person/{id}", http.HandlerFunc(GetById))

	fmt.Printf("server running on port :3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
