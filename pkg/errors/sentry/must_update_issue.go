package sentry

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) MustUpdateIssue(
	organization string,
	identifier string,
	status string,
	assignedTo string,
) *response.Issue {
	result, e := c.UpdateIssue(organization, identifier, status, assignedTo)
	errors.PanicOnError(e)

	return result
}
