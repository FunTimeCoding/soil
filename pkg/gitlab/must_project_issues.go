package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/issue"
)

func (c *Client) MustProjectIssues(project int64) []*issue.Issue {
	result, e := c.ProjectIssues(project)
	errors.PanicOnError(e)

	return result
}
