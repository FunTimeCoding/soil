package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) Organization(slug string) (*response.Organization, error) {
	b, e := c.basic.Get(
		fmt.Sprintf("organizations/%s", slug),
		nil,
	)

	if e != nil {
		return nil, e
	}

	var result response.Organization

	if f := json.Unmarshal(b, &result); f != nil {
		return nil, f
	}

	return &result, nil
}
