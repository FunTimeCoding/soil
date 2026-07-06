package sentry

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) MustLatestEvent(
	organization string,
	issueIdentifier string,
) *response.Event {
	result, e := c.LatestEvent(organization, issueIdentifier)
	errors.PanicOnError(e)

	return result
}
