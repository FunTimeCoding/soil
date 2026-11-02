package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustDeletePipeline(
	x context.Context,
	project int64,
	pipeline int64,
) {
	errors.PanicOnError(c.DeletePipeline(x, project, pipeline))
}
