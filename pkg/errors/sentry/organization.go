package sentry

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) Organization(slug string) (*response.Organization, error) {
	var result response.Organization

	if e := c.basic.Get(fmt.Sprintf("organizations/%s", slug), nil, &result); e != nil {
		return nil, e
	}

	return &result, nil
}
