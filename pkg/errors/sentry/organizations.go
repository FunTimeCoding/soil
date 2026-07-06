package sentry

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) Organizations() ([]response.Organization, error) {
	b, e := c.basic.Get("organizations", nil)

	if e != nil {
		return nil, e
	}

	var result []response.Organization

	if f := json.Unmarshal(b, &result); f != nil {
		return nil, f
	}

	return result, nil
}
