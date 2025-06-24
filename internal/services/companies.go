// Company business logic
package services

import (
	"logistics-crm/internal/models"
	"logistics-crm/internal/database"
	"logistics-crm/internal/integrations/apollo"
)

func EnrichCompany(domain string, apolloClient *apollo.Client) (*models.Company, error) {
	// Check if company exists
	existing, err := database.GetCompanyByDomain(domain)
	if err == nil {
		return existing, nil
	}

	// Fetch from Apollo.io
	apolloData, err := apolloClient.GetCompanyProfile(domain)
	if err != nil {
		return nil, err
	}

	// Convert and save
	company := &models.Company{
		Domain:   domain,
		Name:     apolloData.Name,
		Industry: apolloData.Industry,
		Revenue:  apolloData.Revenue,
	}

	// TODO: save to db
}
