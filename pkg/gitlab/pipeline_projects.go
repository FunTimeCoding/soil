package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	library "github.com/funtimecoding/soil/pkg/project"
)

func (c *Client) PipelineProjects() []*project.Project {
	return c.ProjectsWithFile(library.GitLabFile, false)
}
