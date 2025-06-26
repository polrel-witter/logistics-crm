// Company CRUD operations
package database

import (
	"log"
	"logistics-crm/internal/models"
)

// Methods (operating on DB)
func (db *DB) UpdateCompany(company *models.Company) error {
	// TODO: update to include full query
	query := `UPDATE companies
            SET name = ?, updated_at = datetime('now')
            WHERE id = ?`

	_, err := db.conn.Exec(query,
		company.Domain,
		company.Name,
		company.CgCode,
		company.Note,
		company.Industry,
		company.Revenue,
		company.ID,
	)

	if err != nil {
		log.Printf("Error updating company: %v", err)
		return err
	}

	log.Printf("Successfully updated company with id: %d", company.ID)
	return nil
}

func (db *DB) CreateCompany(company *models.Company) error {
	query := `INSERT INTO companies (domain, name, cg_code, note, industry, revenue)
			  VALUES (?, ?, ?, ?, ?, ?) 
			  RETURNING id, created_at, updated_at`

	err := db.conn.QueryRow(query,
		company.Domain,
		company.Name,
		company.CgCode,
		company.Note,
		company.Industry,
		company.Revenue,
	).Scan(&company.ID, &company.CreatedAt, &company.UpdatedAt)

	if err != nil {
		log.Printf("Error creating company: %v", err)
		return err
	}

	log.Printf("Successfully created company with id: %d", company.ID)
	return nil
}

func (db *DB) GetAllCompanies() ([]*models.Company, error) {
	query := `SELECT id, domain, name, cg_code, note, industry, revenue, created_at, updated_at
			  FROM companies
			  ORDER BY created_at DESC`

	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []*models.Company
	for rows.Next() {
		company := &models.Company{}
		err := rows.Scan(company.ScanFields()...)
		if err != nil {
			return nil, err
		}
		log.Println(company)
		companies = append(companies, company)
	}

	return companies, nil
}

func (db *DB) GetCompanyByDomain(domain string) (*models.Company, error) {
	query := `SELECT id, domain, name, cg_code, note, industry, revenue, created_at, updated_at
			  FROM companies`

	company := &models.Company{}
	err := db.conn.QueryRow(query, domain).Scan(company.ScanFields()...)
	return company, err
}
