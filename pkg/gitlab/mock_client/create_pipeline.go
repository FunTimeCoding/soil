package mock_client

import (
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline_detail"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) CreatePipeline(
	_ int64,
	_ string,
	_ []*gitlab.PipelineVariableOptions,
) (*pipeline_detail.Detail, error) {
	return nil, nil
}
