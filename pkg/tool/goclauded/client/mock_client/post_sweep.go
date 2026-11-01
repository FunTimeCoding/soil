package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) PostSweep(_ context.Context) (*client.SweepResponse, error) {
	return &client.SweepResponse{}, c.Failure
}
