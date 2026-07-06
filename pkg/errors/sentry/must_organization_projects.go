package sentry

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) MustOrganizationProjects(organization string) []response.Project {
	result, e := c.OrganizationProjects(organization)
	errors.PanicOnError(e)

	return result
}
