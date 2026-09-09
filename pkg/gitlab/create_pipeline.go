package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline_detail"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) CreatePipeline(
	project int64,
	reference string,
	v []*gitlab.PipelineVariableOptions,
) (*pipeline_detail.Detail, error) {
	result, _, e := c.client.Pipelines.CreatePipeline(
		project,
		&gitlab.CreatePipelineOptions{Ref: new(reference), Variables: &v},
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return pipeline_detail.New(result), nil
}
