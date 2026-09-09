package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline_detail"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MustCreatePipeline(
	project int64,
	reference string,
	v []*gitlab.PipelineVariableOptions,
) *pipeline_detail.Detail {
	result, e := c.CreatePipeline(project, reference, v)
	errors.PanicOnError(e)

	return result
}
