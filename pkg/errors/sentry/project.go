package sentry

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) Project(
	organization string,
	projectSlug string,
) (*response.Project, error) {
	var result response.Project

	if e := c.basic.Get(
		fmt.Sprintf("projects/%s/%s", organization, projectSlug),
		nil,
		&result,
	); e != nil {
		return nil, e
	}

	return &result, nil
}
