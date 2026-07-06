package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) Pipelines(project int64) ([]*gitlab.PipelineInfo, error) {
	result, _, e := c.client.Pipelines.ListProjectPipelines(
		project,
		&gitlab.ListProjectPipelinesOptions{
			ListOptions: gitlab.ListOptions{PerPage: constant.PerPage1000},
		},
	)

	return result, e
}
