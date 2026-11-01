package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetUsage(_ context.Context) (*client.UsageResponse, bool, error) {
	return &client.UsageResponse{}, true, c.Failure
}
