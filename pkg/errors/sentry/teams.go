package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) Teams(organization string) ([]response.Team, error) {
	b, e := c.basic.Get(
		fmt.Sprintf("organizations/%s/teams", organization),
		nil,
	)

	if e != nil {
		return nil, e
	}

	var result []response.Team

	if f := json.Unmarshal(b, &result); f != nil {
		return nil, f
	}

	return result, nil
}
