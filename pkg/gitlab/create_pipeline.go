package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) CreatePipeline(
	project int64,
	reference string,
	v []*gitlab.PipelineVariableOptions,
) (*gitlab.Pipeline, error) {
	result, _, e := c.client.Pipelines.CreatePipeline(
		project,
		&gitlab.CreatePipelineOptions{Ref: new(reference), Variables: &v},
	)

	return result, wrapError(e)
}
