package main

import (
	"html/template"
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

	// Parse templates
	tmpl, err := template.ParseGlob("web/templates/*.html")
	if err != nil {
		log.Fatal("Failed to parse templates:", err)
	}

	// Setup handlers
	companyHandler := handlers.NewCompanyHandler(db, tmpl)

	// Routes
	http.HandleFunc("/companies", companyHandler.ListCompanies)
	http.HandleFunc("/api/companies/create", companyHandler.CreateCompany)

	// Serve static files (for your HTML/CSS/JS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))

	log.Println("Server starting on :8080")
	log.Println("Database: crm.db")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
