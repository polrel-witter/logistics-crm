// Companies handler
package handlers

import (
	"html/template"
	"logistics-crm/internal/database"
	"logistics-crm/internal/integrations/apollo"
	"logistics-crm/internal/models"
	"logistics-crm/internal/services"
	"net/http"
)

type CompanyHandler struct {
	db     *database.DB
	tmpl   *template.Template
	apolloClient *apollo.Client
}

func NewCompanyHandler(db *database.DB, tmpl *template.Template, apolloClient *apollo.Client) *CompanyHandler {
	return &CompanyHandler{
		db:           db,
		tmpl:         tmpl,
		apolloClient: apolloClient,
	}
}

func (h *CompanyHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	domain := r.FormValue("domain")

	company := &models.Company{
		Domain: domain,
	}

	// Save to database
	if err := h.db.CreateCompany(company); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// Update database with info pulled from Apollo.io
	services.EnrichCompany(domain, h.apolloClient)

	// For HTMX: return just the new company card
	w.Header().Set("Content-Type", "text/html")
	if err := h.tmpl.ExecuteTemplate(w, "company_card.html", company); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *CompanyHandler) ListCompanies(w http.ResponseWriter, r *http.Request) {
	// 1. Get data from database
	companies, err := h.db.GetAllCompanies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 2. Prepare data for template
	data := struct {
		Companies []*models.Company
		Title     string
	}{
		Companies: companies,
		Title:     "Companies",
	}

	// 3. Render template with data
	w.Header().Set("Content-Type", "text/html")
	if err := h.tmpl.ExecuteTemplate(w, "companies_list.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
