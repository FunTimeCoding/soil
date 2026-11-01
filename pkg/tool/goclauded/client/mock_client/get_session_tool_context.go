package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetSessionToolContext(_ context.Context, _ string, _ *client.GetSessionToolContextParams) (*client.ToolContextResponse, error) {
	return &client.ToolContextResponse{}, c.Failure
}
