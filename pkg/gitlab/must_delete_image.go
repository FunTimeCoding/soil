package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustDeleteImage(
	x context.Context,
	project int64,
	repository int64,
	tag string,
) {
	errors.PanicOnError(c.DeleteImage(x, project, repository, tag))
}
