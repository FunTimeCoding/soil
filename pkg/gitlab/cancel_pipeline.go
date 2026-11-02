package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline_detail"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) CancelPipeline(
	x context.Context,
	project int64,
	identifier int64,
) (*pipeline_detail.Detail, error) {
	result, _, e := c.client.Pipelines.CancelPipelineBuild(
		project,
		identifier,
		gitlab.WithContext(x),
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return pipeline_detail.New(result), nil
}
