package sentry

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) Projects() ([]response.Project, error) {
	b, e := c.basic.Get("projects", nil)

	if e != nil {
		return nil, e
	}

	var result []response.Project

	if f := json.Unmarshal(b, &result); f != nil {
		return nil, f
	}

	return result, nil
}
