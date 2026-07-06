package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) MustProject(identifier int64) *project.Project {
	result, e := c.Project(identifier)
	errors.PanicOnError(e)

	return result
}
