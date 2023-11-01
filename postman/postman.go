package postman

import (
	"fmt"
	"strings"
)

type Postman struct {
	ApiKey       string `json:"api_key"`
	WorkspaceID  string `json:"workspace_id"`
	CollectionID string `json:"collection_id"`
}

func (p *Postman) validation() error {
	if p == nil {
		return fmt.Errorf("empty values")
	}

	if strings.TrimSpace(p.ApiKey) == "" {
		return fmt.Errorf("api key cannot be empty")
	}

	if strings.TrimSpace(p.WorkspaceID) == "" {
		return fmt.Errorf("workspace id cannot be empty")
	}

	if strings.TrimSpace(p.CollectionID) == "" {
		return fmt.Errorf("collection id cannot be empty")
	}

	return nil
}
