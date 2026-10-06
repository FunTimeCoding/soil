package sentry

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) OrganizationProjects(organization string) ([]response.Project, error) {
	var result []response.Project

	if e := c.basic.Get(
		fmt.Sprintf("organizations/%s/projects", organization),
		nil,
		&result,
	); e != nil {
		return nil, e
	}

	return result, nil
}
