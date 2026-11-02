package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustDeletePipeline(
	x context.Context,
	project int64,
	identifier int64,
) {
	errors.PanicOnError(c.DeletePipeline(x, project, identifier))
}
