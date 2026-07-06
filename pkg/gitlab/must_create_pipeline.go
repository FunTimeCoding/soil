package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MustCreatePipeline(
	project int64,
	reference string,
	v []*gitlab.PipelineVariableOptions,
) *gitlab.Pipeline {
	result, e := c.CreatePipeline(project, reference, v)
	errors.PanicOnError(e)

	return result
}
