package postman

import (
	"fmt"
	"time"
)

type GetAPIVersionsResponse struct {
	Versions []struct {
		ID                    string      `json:"id"`
		Name                  string      `json:"name"`
		Summary               interface{} `json:"summary"`
		CreatedBy             string      `json:"createdBy"`
		UpdatedBy             string      `json:"updatedBy"`
		Stage                 string      `json:"stage"`
		Visibility            string      `json:"visibility"`
		Api                   string      `json:"api"`
		CreatedAt             time.Time   `json:"createdAt"`
		UpdatedAt             time.Time   `json:"updatedAt"`
		RepositoryIntegration interface{} `json:"repositoryIntegration"`
	} `json:"versions"`
}

func (p *Postman) GetAPIVersions(apiID string) (*GetAPIVersionsResponse, error) {
	if err := p.validation(); err != nil {
		return nil, err
	}

	var (
		url      = fmt.Sprintf("https://api.getpostman.com/apis/%s/versions", apiID)
		method   = "GET"
		payload  interface{}
		response *GetAPIVersionsResponse
	)

	err := p.DoRequestAndUnmarshal(method, url, payload, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
