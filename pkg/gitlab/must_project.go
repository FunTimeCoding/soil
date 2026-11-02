package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) MustProject(
	x context.Context,
	identifier int64,
) *project.Project {
	result, e := c.Project(x, identifier)
	errors.PanicOnError(e)

	return result
}
