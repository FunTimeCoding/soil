package sentry

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) Whoami() (*response.User, error) {
	b, e := c.basic.Get("auth", nil)

	if e != nil {
		return nil, e
	}

	var result response.User

	if f := json.Unmarshal(b, &result); f != nil {
		return nil, f
	}

	return &result, nil
}
