package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetSessionWindow(_ context.Context, _ string, _ *client.GetSessionWindowParams) (*client.WindowResponse, error) {
	return &client.WindowResponse{}, c.Failure
}
