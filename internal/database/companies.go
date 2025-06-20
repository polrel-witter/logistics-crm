// Company database actions
package database

import (
	"logistics-crm/internal/models"
)

// TODO: alias the db query to reduce clutter

func (db *DB) Save(company *models.Company) error {
	if company.ID == 0 {
		return db.CreateCompany(company)
	}
	return db.UpdateCompany(company)
}

func (db *DB) CreateCompany(company *models.Company) error {
	query := `
        INSERT INTO companies (domain, name, cg_code, note, revenue, locations, industry)
        VALUES (?, ?, ?, ?, datetime('now'))
        RETURNING id, created_at
    `
	result, err := db.conn.Exec(query, company.Name, company.Domain,
		company.CgCode, company.Note, company.Revenue, company.Locations, company.Industry)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	company.ID = int(id)
	return nil
}

func (db *DB) UpdateCompany(company *models.Company) error {
	query := `
        UPDATE companies 
        SET name = ?, domain = ?, cg_code = ?, note = ?, revenue = ?, locations = ?, industry = ?
        WHERE id = ?
    `
	_, err := db.conn.Exec(query, company.Name, company.Domain,
		company.CgCode, company.Note, company.Revenue, company.Locations, company.Industry)
	return err
}

func (db *DB) GetAllCompanies() ([]*models.Company, error) {
	query := `SELECT id, domain, name, revenue, industry, created_at FROM companies ORDER BY created_at DESC`

	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []*models.Company
	for rows.Next() {
		company := &models.Company{}
		err := rows.Scan(&company.ID, &company.Name, &company.Domain,
			&company.CgCode, &company.Note, &company.Revenue, &company.Locations,
			&company.Industry, &company.CreatedAt, &company.UpdatedAt)
		if err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}

	return companies, nil
}

func (db *DB) GetCompanyByDomain(domain string) (*models.Company, error) {
	query := `SELECT id, domain, name, revenue, industry, created_at FROM companies WHERE domain = ?`

	company := &models.Company{}
	err := db.conn.QueryRow(query, domain).Scan(
		&company.ID, &company.Name, &company.Domain,
		&company.CgCode, &company.Note, &company.Revenue, &company.Locations,
		&company.Industry, &company.CreatedAt, &company.UpdatedAt)

	return company, err
}
