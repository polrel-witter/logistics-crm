// Company database actions
package database

import (
	"log"
	"logistics-crm/internal/models"
)

// TODO: add update capability
//func (db *DB) SaveCompany(company *models.Company) error {
//	if company.ID == 0 {
//		db.createCompany(company)
//		return nil
//	}
//
//	db.updateCompany(company)
//	return nil
//}

//func (db *DB) updateCompany(company *models.Company) error {
//	query := `UPDATE companies
//            SET name = ?, updated_at = datetime('now')
//            WHERE id = ?`
//
//	_, err := db.conn.Exec(query,
//		company.Domain,
//		company.Name,
//		company.CgCode,
//		company.Note,
//		company.Industry,
//		company.Revenue,
//		company.ID,
//	)
//
//	return err
//}

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

	log.Printf("Successfully added company with id: %d", company.ID)
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

// TODO: domain search
// func (db *DB) GetCompanyByDomain(domain string) (*models.Company, error) {
// 	query := `SELECT id, domain, name, cg_code, note, industry, revenue, created_at, updated_at
// 			  FROM companies`
//
// 	company := &models.Company{}
// 	err := db.conn.QueryRow(query, domain).Scan(company.ScanFields()...)
// 	return company, err
// }
