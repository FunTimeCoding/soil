package sentry

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) MustOrganizations() []response.Organization {
	result, e := c.Organizations()
	errors.PanicOnError(e)

	return result
}
