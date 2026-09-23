package mock_client

import (
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) SeedPipeline(
	identifier int64,
	reference string,
	hash string,
	status string,
) {
	c.pipelines = append(
		c.pipelines,
		pipeline.New(
			&gitlab.PipelineInfo{
				ID:     identifier,
				Ref:    reference,
				SHA:    hash,
				Status: status,
			},
		),
	)
}
