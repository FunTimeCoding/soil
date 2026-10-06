package sentry

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (c *Client) Projects() ([]response.Project, error) {
	var result []response.Project

	if e := c.basic.Get("projects", nil, &result); e != nil {
		return nil, e
	}

	return result, nil
}
