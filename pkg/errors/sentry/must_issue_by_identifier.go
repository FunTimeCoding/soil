package sentry

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) MustIssueByIdentifier(
	organization string,
	identifier string,
) *response.Issue {
	result, e := c.IssueByIdentifier(organization, identifier)
	errors.PanicOnError(e)

	return result
}
