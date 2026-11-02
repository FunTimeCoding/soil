package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustDeletePackage(
	x context.Context,
	project int64,
	repository int64,
) {
	errors.PanicOnError(c.DeletePackage(x, project, repository))
}
