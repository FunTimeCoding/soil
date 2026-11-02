package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) Project(
	_ context.Context,
	_ int64,
) (*project.Project, error) {
	return nil, nil
}
