package sentry

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) MustOrganization(slug string) *response.Organization {
	result, e := c.Organization(slug)
	errors.PanicOnError(e)

	return result
}
