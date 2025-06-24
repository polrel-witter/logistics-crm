package main

import (
	"html/template"
	"log"
	"logistics-crm/internal/database"
	"logistics-crm/internal/handlers"
	"logistics-crm/internal/integrations/apollo"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	// Initialize database
	db, err := database.New()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Parse html templates
	tmpl, err := template.ParseGlob("web/templates/*.html")
	if err != nil {
		log.Fatal("Failed to parse templates:", err)
	}

	// Configure integrations
	apolloAPIKey := os.Getenv("APOLLO_API_KEY")
	if apolloAPIKey == "" {
		log.Fatal("APOLLO_API_KEY environment variable is required")
	}
	apolloClient := &apollo.Client{
		APIKey: apolloAPIKey,
		Client: &http.Client{},
	}

	// Setup handlers
	companyHandler := handlers.NewCompanyHandler(db, tmpl, apolloClient)

	// Routes
	http.HandleFunc("/companies", companyHandler.ListCompanies)
	http.HandleFunc("/api/companies/create", companyHandler.CreateCompany)

	// Serve static files (for your HTML/CSS/JS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))

	log.Println("Server starting on :8080")
	log.Println("Database: crm.db")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
