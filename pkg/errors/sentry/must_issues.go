package sentry

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/issue"
)

func (c *Client) MustIssues(
	organization string,
	projectIdentifier string,
	period string,
) []*issue.Issue {
	result, e := c.Issues(organization, projectIdentifier, period)
	errors.PanicOnError(e)

	return result
}
