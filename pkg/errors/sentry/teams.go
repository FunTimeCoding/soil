package sentry

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) Teams(organization string) ([]response.Team, error) {
	var result []response.Team

	if e := c.basic.Get(
		fmt.Sprintf("organizations/%s/teams", organization),
		nil,
		&result,
	); e != nil {
		return nil, e
	}

	return result, nil
}
