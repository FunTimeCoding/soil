package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) MustProjects(x context.Context) []*project.Project {
	result, e := c.Projects(x)
	errors.PanicOnError(e)

	return result
}
