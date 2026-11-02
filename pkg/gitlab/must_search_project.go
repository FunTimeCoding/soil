package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) MustSearchProject(
	x context.Context,
	query string,
) []*project.Project {
	result, e := c.SearchProject(x, query)
	errors.PanicOnError(e)

	return result
}
