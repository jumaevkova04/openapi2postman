package postman

import (
	"encoding/json"
	"fmt"
)

func (p *Postman) GetCollection(collectionID string) ([]byte, error) {
	if err := p.validation(); err != nil {
		return nil, err
	}

	var (
		method   = "GET"
		url      = fmt.Sprintf("https://api.getpostman.com/collections/%s", collectionID)
		payload  interface{}
		response interface{}
	)

	err := p.DoRequestAndUnmarshal(method, url, payload, &response)
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(&response)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// ---------------------------------------------------------------------------------------------------------------------

type ReplaceCollectionsDataResponse struct {
	Collection struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Uid  string `json:"uid"`
	} `json:"collection"`
}

func (p *Postman) ReplaceCollectionsData(collectionID string, body []byte) (*ReplaceCollectionsDataResponse, error) {
	if err := p.validation(); err != nil {
		return nil, err
	}

	var (
		method   = "PUT"
		url      = fmt.Sprintf("https://api.getpostman.com/collections/%s", collectionID)
		payload  = body
		response *ReplaceCollectionsDataResponse
	)

	err := p.DoRequestAndUnmarshal(method, url, payload, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ---------------------------------------------------------------------------------------------------------------------

type UpdateCollectionsDataResponse struct {
	Collection struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Uid  string `json:"uid"`
	} `json:"collection"`
}

func (p *Postman) UpdateCollectionsData(collectionID string, body []byte) (*UpdateCollectionsDataResponse, error) {
	if err := p.validation(); err != nil {
		return nil, err
	}

	var (
		method   = "PATCH"
		url      = fmt.Sprintf("https://api.getpostman.com/collections/%s", collectionID)
		payload  = body
		response *UpdateCollectionsDataResponse
	)

	err := p.DoRequestAndUnmarshal(method, url, payload, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ---------------------------------------------------------------------------------------------------------------------

type DeleteCollectionResponse struct {
	Collection struct {
		ID  string `json:"id"`
		UID string `json:"uid"`
	} `json:"collection"`
}

func (p *Postman) DeleteCollection(collectionID string) (*DeleteCollectionResponse, error) {
	if err := p.validation(); err != nil {
		return nil, err
	}

	var (
		method   = "DELETE"
		url      = fmt.Sprintf("https://api.getpostman.com/collections/%s", collectionID)
		payload  interface{}
		response *DeleteCollectionResponse
	)

	err := p.DoRequestAndUnmarshal(method, url, payload, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
