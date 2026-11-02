package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline_detail"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) CreatePipeline(
	_ context.Context,
	_ int64,
	_ string,
	_ []*gitlab.PipelineVariableOptions,
) (*pipeline_detail.Detail, error) {
	return nil, c.pipelineFailure
}
