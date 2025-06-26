// Companies handler
package handlers

import (
	"html/template"
	"logistics-crm/internal/database"
	"logistics-crm/internal/models"
	"logistics-crm/internal/services"
	"net/http"
)

// Types
type CompanyHandler struct {
	db       *database.DB
	services *services.CompanyService
	tmpl     *template.Template
}

// Error handling
func (h *CompanyHandler) writeInternalError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func (h *CompanyHandler) writeMethodNotAllowed(w http.ResponseWriter) {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// Methods
func NewCompanyHandler(db *database.DB, services *services.CompanyService, tmpl *template.Template) *CompanyHandler {
	return &CompanyHandler{
		db:       db,
		services: services,
		tmpl:     tmpl,
	}
}

func (h *CompanyHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.writeMethodNotAllowed(w)
		return
	}

	domain := r.FormValue("domain")

	company := &models.Company{
		Domain: domain,
	}

	// Save to database
	if err := h.db.CreateCompany(company); err != nil {
		h.writeInternalError(w, err)
		return
	}

	// Update database with info pulled from Apollo.io
	if err := h.services.EnrichCompany(domain); err != nil {
		h.writeInternalError(w, err)
		return
	}

	// For HTMX: return just the new company card
	w.Header().Set("Content-Type", "text/html")
	if err := h.tmpl.ExecuteTemplate(w, "company_card.html", company); err != nil {
		h.writeInternalError(w, err)
		return
	}
}

func (h *CompanyHandler) ListCompanies(w http.ResponseWriter, r *http.Request) {
	// Get data from database
	companies, err := h.db.GetAllCompanies()
	if err != nil {
		h.writeInternalError(w, err)
		return
	}

	// Prepare data for template
	data := struct {
		Companies []*models.Company
		Title     string
	}{
		Companies: companies,
		Title:     "Companies",
	}

	// Render template with data
	w.Header().Set("Content-Type", "text/html")
	if err := h.tmpl.ExecuteTemplate(w, "companies_list.html", data); err != nil {
		h.writeInternalError(w, err)
		return
	}
}
