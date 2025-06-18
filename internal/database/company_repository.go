// Company database actions
package database

import (
	"database/sql"
	"logistics-crm/internal/models"
)

type CompanyRepository struct {
	db *sql.DB
}

// These are your "actions" - but separated from the model
func (r *CompanyRepository) Save(company *models.Company) error {
	if company.ID == 0 {
		return r.create(company)
	}
	return r.update(company)
}

func (r *CompanyRepository) create(company *models.Company) error {
	query := `
        INSERT INTO companies (domain, name, cg_code, usher, note, revenue, locations, industry)
        VALUES (?, ?, ?, ?, datetime('now'))
        RETURNING id, created_at
    `
	return r.db.QueryRow(query,
		company.Domain, company.Name, company.CgCode, company.Note, company.Revenue, company.Revenue, company.Industry,
	).Scan(&company.ID, &company.CreatedAt)
}

func (r *CompanyRepository) update(company *models.Company) error {
	query := `
        UPDATE companies 
        SET name = ?, revenue = ?, industry = ?
        WHERE id = ?
    `
	_, err := r.db.Exec(query, company.Name, company.Revenue, company.Industry, company.ID)
	return err
}

// func (r *CompanyRepository) GetByDomain(domain string) (*models.Company, error) {
// 	// Implementation...
// }
//
// func (r *CompanyRepository) GetByID(id int) (*models.Company, error) {
// 	// Implementation...
// }
