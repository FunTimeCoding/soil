package sentry

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/issue"
)

func (c *Client) MustIssuesSimple(verbose bool) []*issue.Issue {
	result, e := c.IssuesSimple(verbose)
	errors.PanicOnError(e)

	return result
}
