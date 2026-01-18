package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /person", http.HandlerFunc(Get))

	fmt.Printf("server running on port :3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
