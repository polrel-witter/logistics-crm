// Company business logic
package services

import (
	"logistics-crm/internal/database"
	"logistics-crm/internal/integrations/apollo"
	"logistics-crm/internal/models"
)

type CompanyService struct {
	db     *database.DB
	apollo *apollo.Client
}

func NewCompanyService(db *database.DB, apollo *apollo.Client) *CompanyService {
	return &CompanyService{
		db:     db,
		apollo: apollo,
	}
}

func (s *CompanyService) EnrichCompany(domain string) error {
	// Fetch from Apollo.io
	apolloData, err := s.apollo.GetCompanyProfile(domain)
	if err != nil {
		return err
	}

	company := &models.Company{
		Domain:   domain,
		Name:     apolloData.Name,
		Industry: apolloData.Industry,
		Revenue:  apolloData.Revenue,
	}

	return s.db.UpdateCompany(company)
}
