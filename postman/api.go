package postman

import (
	"fmt"
	"time"
)

type CreateAPIResponse struct {
	Api struct {
		ID          string    `json:"id"`
		Name        string    `json:"name"`
		Summary     string    `json:"summary"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
		CreatedBy   string    `json:"createdBy"`
		UpdatedBy   string    `json:"updatedBy"`
	} `json:"api"`
}

type CreateAPIRequest struct {
	Api Api `json:"api"`
}

type Api struct {
	Name        string `json:"name"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
}

func (p *Postman) CreateAPI(request CreateAPIRequest) (*CreateAPIResponse, error) {
	if err := p.validation(); err != nil {
		return nil, err
	}

	var (
		method   = "POST"
		url      = "https://api.getpostman.com/apis"
		payload  = request
		response *CreateAPIResponse
	)

	err := p.DoRequestAndUnmarshal(method, url, payload, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ---------------------------------------------------------------------------------------------------------------------

type DeleteAPIResponse struct {
	Api struct {
		ID string `json:"id"`
	} `json:"api"`
}

func (p *Postman) DeleteAPI(apiID string) (*DeleteAPIResponse, error) {
	if err := p.validation(); err != nil {
		return nil, err
	}

	var (
		method   = "DELETE"
		url      = fmt.Sprintf("https://api.getpostman.com/apis/%s", apiID)
		payload  interface{}
		response *DeleteAPIResponse
	)

	err := p.DoRequestAndUnmarshal(method, url, payload, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
