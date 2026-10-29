package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
)

func (c *Client) ListProjects(_ context.Context) ([]*client.JiraProject, error) {
	return nil, c.Failure
}
