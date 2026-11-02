package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/file"
)

func (c *Client) MustFile(
	x context.Context,
	project int64,
	branch string,
	name string,
) *file.File {
	result, e := c.File(x, project, branch, name)
	errors.PanicOnError(e)

	return result
}
