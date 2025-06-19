package main

import (
	"log"
	"logistics-crm/internal/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/companies", handlers.CompaniesHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
