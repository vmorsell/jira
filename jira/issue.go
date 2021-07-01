package jira

import (
	"encoding/json"
	"fmt"
)

type Issue struct {
	ID   string `json:"id"`
	Self string `json:"self"`
	Key  string `json:"key"`
}

func (c *Client) Issue(idOrKey string) (*Issue, error) {
	url := fmt.Sprintf("/issue/%s", idOrKey)
	res, err := c.Req("GET", url)
	if err != nil {
		return nil, fmt.Errorf("issues: %w", err)
	}

	var out *Issue
	if err := json.Unmarshal(res, &out); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}
	return out, nil
}
