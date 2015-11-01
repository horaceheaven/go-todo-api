package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting API")
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
