package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) MustProjects() []*project.Project {
	result, e := c.Projects()
	errors.PanicOnError(e)

	return result
}
