package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetCheck(_ context.Context, _ *client.GetCheckParams) (*client.CheckResponse, error) {
	return &client.CheckResponse{}, c.Failure
}
