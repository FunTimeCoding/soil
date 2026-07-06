package sentry

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/issue"
)

func (c *Client) MustAllIssues() []*issue.Issue {
	result, e := c.AllIssues()
	errors.PanicOnError(e)

	return result
}
