package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) ProjectByName(
	_ context.Context,
	_ string,
	_ string,
) (*project.Project, error) {
	return nil, c.failure
}
