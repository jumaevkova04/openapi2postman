package postman

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (p *Postman) DoRequestAndUnmarshal(method string, url string, payload interface{}, response interface{}) error {
	req, err := p.createRequest(method, url, payload)
	if err != nil {
		return err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent &&
		resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("code: %d; err: %+v", resp.StatusCode, string(body))
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postman) createRequest(method string, url string, payload interface{}) (*http.Request, error) {
	var data []byte
	var err error

	if payload != nil {
		_, isPayloadSliceByte := payload.([]byte)
		if isPayloadSliceByte {
			data = payload.([]byte)
		} else {
			data, err = json.Marshal(payload)
			if err != nil {
				return nil, fmt.Errorf("json marshal %s: %w", payload, err)
			}
		}
	}

	request, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("new http request: %w", err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", p.ApiKey)

	return request, nil
}
