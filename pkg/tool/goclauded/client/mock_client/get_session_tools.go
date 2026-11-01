package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetSessionTools(_ context.Context, _ string) (*client.ToolsResponse, error) {
	return &client.ToolsResponse{}, c.Failure
}
