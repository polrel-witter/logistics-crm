package main

import (
	"log"
	"logistics-crm/internal/database"
	"logistics-crm/internal/handlers"
	"net/http"
)

func main() {
	// Initialize database
	db, err := database.New()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Setup handlers
	companyHandler := handlers.NewCompanyHandler(db)

	// Routes
	http.HandleFunc("/api/companies/create", companyHandler.CreateCompany)

	// Serve static files (for your HTML/CSS/JS)
	http.Handle("/", http.FileServer(http.Dir("./web/")))

	log.Println("Server starting on :8080")
	log.Println("Database: crm.db")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
