package github

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/github/issue"
)

func (c *Client) MustIssues() []*issue.Issue {
	result, e := c.Issues()
	errors.PanicOnError(e)

	return result
}
