package gitlab

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) PipelineProjects(
	x context.Context,
) ([]*project.Project, error) {
	return c.ProjectsWithFile(x, soil.GitLabFile, false)
}
