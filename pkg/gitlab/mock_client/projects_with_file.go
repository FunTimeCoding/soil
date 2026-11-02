package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) ProjectsWithFile(
	_ context.Context,
	_ string,
	_ bool,
) ([]*project.Project, error) {
	return nil, c.failure
}
