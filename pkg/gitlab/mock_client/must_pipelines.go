package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
)

func (c *Client) MustPipelines(
	_ context.Context,
	_ int64,
) []*pipeline.Pipeline {
	return c.pipelines
}
