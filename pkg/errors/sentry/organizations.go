package sentry

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (c *Client) Organizations() ([]response.Organization, error) {
	var result []response.Organization

	if e := c.basic.Get("organizations", nil, &result); e != nil {
		return nil, e
	}

	return result, nil
}
