package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) OrganizationProjects(organization string) ([]response.Project, error) {
	b, e := c.basic.Get(
		fmt.Sprintf("organizations/%s/projects", organization),
		nil,
	)

	if e != nil {
		return nil, e
	}

	var result []response.Project

	if f := json.Unmarshal(b, &result); f != nil {
		return nil, f
	}

	return result, nil
}
