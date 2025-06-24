package apollo

import (
	"net/http"
)

type Client struct {
	APIKey string
	Client *http.Client
}

type CompanyProfile struct {
	Name     string `json:"name"`
	Revenue  int64  `json:"annual_revenue"`
	Industry string `json:"industry"`
}

func (c *Client) GetCompanyProfile(domain string) (*CompanyProfile, error) {
	// TODO: API call implementation
	return &CompanyProfile{}, nil
}
