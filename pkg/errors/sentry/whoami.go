package sentry

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (c *Client) Whoami() (*response.User, error) {
	var result response.User

	if e := c.basic.Get("auth", nil, &result); e != nil {
		return nil, e
	}

	return &result, nil
}
