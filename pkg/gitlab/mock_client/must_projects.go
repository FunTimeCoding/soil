package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) MustProjects(_ context.Context) []*project.Project {
	return c.projects
}
