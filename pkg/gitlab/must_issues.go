package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/issue"
)

func (c *Client) MustIssues() []*issue.Issue {
	result, e := c.Issues()
	errors.PanicOnError(e)

	return result
}
