package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MustPipelines(project int64) []*gitlab.PipelineInfo {
	result, e := c.Pipelines(project)
	errors.PanicOnError(e)

	return result
}
