package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
)

func (c *Client) MustPipelines(
	x context.Context,
	project int64,
) []*pipeline.Pipeline {
	result, e := c.Pipelines(x, project, "", "", 0)
	errors.PanicOnError(e)

	return result
}
