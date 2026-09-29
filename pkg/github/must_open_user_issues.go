package github

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/github/issue"
)

func (c *Client) MustOpenUserIssues() []*issue.Issue {
	result, e := c.OpenUserIssues()
	errors.PanicOnError(e)

	return result
}
