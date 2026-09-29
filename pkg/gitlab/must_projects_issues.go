package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/issue"
)

func (c *Client) MustProjectsIssues() []*issue.Issue {
	result, e := c.ProjectsIssues()
	errors.PanicOnError(e)

	return result
}
