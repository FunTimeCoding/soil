package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) MustPipelineProjects() []*project.Project {
	result, e := c.PipelineProjects()
	errors.PanicOnError(e)

	return result
}
