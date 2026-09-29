package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) MustProjectsWithFile(
	path string,
	caseInsensitive bool,
) []*project.Project {
	result, e := c.ProjectsWithFile(path, caseInsensitive)
	errors.PanicOnError(e)

	return result
}
