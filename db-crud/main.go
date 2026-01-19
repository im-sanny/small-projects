package main

import (
	"db-crud/db"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db.InitDB()
	mux := http.NewServeMux()

	fmt.Printf("server running on port :3000")
	err1 := http.ListenAndServe(":3000", mux)
	if err1 != nil {
		log.Fatal(err1)
	}
}
