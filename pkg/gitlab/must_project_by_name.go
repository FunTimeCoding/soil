package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) MustProjectByName(
	x context.Context,
	namespace string,
	name string,
) *project.Project {
	result, e := c.ProjectByName(x, namespace, name)
	errors.PanicOnError(e)

	return result
}
