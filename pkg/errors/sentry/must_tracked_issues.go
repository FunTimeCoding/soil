package sentry

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/issue"
)

func (c *Client) MustTrackedIssues() []*issue.Issue {
	result, e := c.TrackedIssues()
	errors.PanicOnError(e)

	return result
}
