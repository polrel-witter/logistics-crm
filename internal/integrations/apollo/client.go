package apollo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Types
type Client struct {
	APIKey string
	Client *http.Client
}

type CompanyProfile struct {
	Name     *string `json:"name"`
	Revenue  *int    `json:"annual_revenue"`
	Industry *string `json:"industry"`
}

// Methods
func (c *Client) buildApolloGETRequest(url string) (*http.Request, error) {
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-api-key", c.APIKey)

	return req, nil
}

func (c *Client) GetCompanyProfile(domain string) (*CompanyProfile, error) {
	url := fmt.Sprintf("https://api.apollo.io/api/v1/organizations/enrich?domain=%s", domain)

	// Http request/response
	req, _ := c.buildApolloGETRequest(url)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Read body
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))

	// Parse body
	var company CompanyProfile
	err = json.Unmarshal(body, &company)
	if err != nil {
		// handle error
		log.Printf("Error parsing JSON: %v", err)
		return nil, err
	}

	return &company, nil
}
