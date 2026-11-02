package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) PipelineProjects(
	_ context.Context,
) ([]*project.Project, error) {
	return nil, c.failure
}
