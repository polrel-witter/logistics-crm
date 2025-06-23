// Companies handler
package handlers

import (
	"html/template"
	"logistics-crm/internal/database"
	"logistics-crm/internal/models"
	"net/http"
)

type CompanyHandler struct {
	db   *database.DB
	tmpl *template.Template
}

func NewCompanyHandler(db *database.DB, tmpl *template.Template) *CompanyHandler {
	return &CompanyHandler{
		db:   db,
		tmpl: tmpl,
	}
}

func (h *CompanyHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	company := &models.Company{
		Domain: r.FormValue("domain"),
		// TODO: pull external info based on domain
	}

	// Save to database
	if err := h.db.CreateCompany(company); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
