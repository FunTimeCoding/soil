package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustDeleteTag(
	x context.Context,
	project int64,
	name string,
) {
	errors.PanicOnError(c.DeleteTag(x, project, name))
}
