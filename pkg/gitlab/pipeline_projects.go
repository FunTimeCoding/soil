package gitlab

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) PipelineProjects() []*project.Project {
	return c.ProjectsWithFile(constant.GitLabFile, false)
}
