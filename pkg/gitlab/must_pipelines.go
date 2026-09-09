package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
)

func (c *Client) MustPipelines(project int64) []*pipeline.Pipeline {
	result, e := c.Pipelines(project, "", "", 0)
	errors.PanicOnError(e)

	return result
}
