// Company database actions
package database

import (
	"database/sql"
	"logistics-crm/internal/models"
)

type CompanyRepository struct {
	db *sql.DB
}

func (r *CompanyRepository) Save(company *models.Company) error {
	if company.ID == 0 {
		return r.Create(company)
	}
	return r.Update(company)
}

func (r *CompanyRepository) Create(company *models.Company) error {
	query := `
        INSERT INTO companies (domain, name, cg_code, note, revenue, locations, industry)
        VALUES (?, ?, ?, ?, datetime('now'))
        RETURNING id, created_at
    `
	return r.db.QueryRow(query,
		company.Domain, company.Name, company.CgCode, company.Note, company.Revenue, company.Revenue, company.Industry,
	).Scan(&company.ID, &company.CreatedAt)
}

func (r *CompanyRepository) Update(company *models.Company) error {
	query := `
        UPDATE companies 
        SET name = ?, cg_code = ?, note = ?, revenue = ?, locations = ?, industry = ?
        WHERE id = ?
    `
	_, err := r.db.Exec(query, company.Name, company.Revenue, company.Industry, company.ID)
	return err
}
