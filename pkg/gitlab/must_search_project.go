package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) MustSearchProject(query string) []*project.Project {
	result, e := c.SearchProject(query)
	errors.PanicOnError(e)

	return result
}
