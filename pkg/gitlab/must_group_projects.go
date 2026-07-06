package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) MustGroupProjects(identifier ...int64) []*project.Project {
	result, e := c.GroupProjects(identifier...)
	errors.PanicOnError(e)

	return result
}
