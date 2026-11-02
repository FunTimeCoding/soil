package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) FileExists(
	_ context.Context,
	_ *project.Project,
	_ string,
	_ string,
) (bool, error) {
	return false, c.failure
}
