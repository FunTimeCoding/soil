package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetChannel(_ context.Context, _ *client.GetChannelParams) (*client.CheckResponse, error) {
	return &client.CheckResponse{}, c.Failure
}
