package sentry

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) MustEvent(
	organization string,
	project string,
	identifier string,
) *response.Event {
	result, e := c.Event(organization, project, identifier)
	errors.PanicOnError(e)

	return result
}
