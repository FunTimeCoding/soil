package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/issue"
)

func (c *Client) MustIssues(x context.Context) []*issue.Issue {
	result, e := c.Issues(x)
	errors.PanicOnError(e)

	return result
}
