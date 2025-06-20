// Companies handler
package handlers

import (
	"logistics-crm/internal/database"
	"logistics-crm/internal/models"
	"net/http"
)

type CompanyHandler struct {
	db *database.DB
}

func NewCompanyHandler(db *database.DB) *CompanyHandler {
	return &CompanyHandler{db: db}
}

func (h *CompanyHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	company := &models.Company{
		Domain: r.FormValue("domain"),
	}

	if err := h.db.CreateCompany(company); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
