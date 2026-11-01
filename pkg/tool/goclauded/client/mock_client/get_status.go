package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetStatus(_ context.Context) (*client.StatusResponse, error) {
	return &client.StatusResponse{}, c.Failure
}
