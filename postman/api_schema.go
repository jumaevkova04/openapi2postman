package postman

import (
	"fmt"
	"time"
)

type CreateAPISchemaResponse struct {
	Schema struct {
		ID         string    `json:"id"`
		Language   string    `json:"language"`
		ApiVersion string    `json:"apiVersion"`
		Type       string    `json:"type"`
		CreatedBy  string    `json:"createdBy"`
		UpdatedBy  string    `json:"updatedBy"`
		CreatedAt  time.Time `json:"createdAt"`
		UpdatedAt  time.Time `json:"updatedAt"`
	} `json:"schema"`
}

type CreateAPISchemaRequest struct {
	Schema Schema `json:"schema"`
}

type Schema struct {
	Language string `json:"language"`
	Schema   string `json:"schema"`
	Type     string `json:"type"`
}

func (p *Postman) CreateAPISchema(apiID string, apiVersionID string, request CreateAPISchemaRequest) (*CreateAPISchemaResponse, error) {
	if err := p.validation(); err != nil {
		return nil, err
	}

	var (
		method   = "POST"
		url      = fmt.Sprintf("https://api.getpostman.com/apis/%s/versions/%s/schemas", apiID, apiVersionID)
		payload  = request
		response *CreateAPISchemaResponse
	)

	err := p.DoRequestAndUnmarshal(method, url, payload, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ---------------------------------------------------------------------------------------------------------------------

type CreateAPICollectionFromSchemaResponse struct {
	Collection struct {
		ID  string `json:"id"`
		UID string `json:"uid"`
	} `json:"collection"`
	Relations []struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	} `json:"relations"`
}

type CreateAPICollectionFromSchemaRequest struct {
	Name      string      `json:"name"`
	Relations []Relations `json:"relations"`
}

type Relations struct {
	Type string `json:"type"`
}

func (p *Postman) CreateAPICollectionFromSchema(apiID string, apiVersionID string, apiSchemaID string, request CreateAPICollectionFromSchemaRequest) (*CreateAPICollectionFromSchemaResponse, error) {
	if err := p.validation(); err != nil {
		return nil, err
	}

	var (
		method   = "POST"
		url      = fmt.Sprintf("https://api.getpostman.com/apis/%s/versions/%s/schemas/%s/collections", apiID, apiVersionID, apiSchemaID)
		payload  = request
		response *CreateAPICollectionFromSchemaResponse
	)

	err := p.DoRequestAndUnmarshal(method, url, payload, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
