package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) MustProjectByName(
	namespace string,
	name string,
) *project.Project {
	result, e := c.ProjectByName(namespace, name)
	errors.PanicOnError(e)

	return result
}
