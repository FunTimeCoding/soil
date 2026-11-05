package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetUsage(_ context.Context) (*client.UsageResponse, bool, error) {
	if c.Failure != nil {
		return nil, false, c.Failure
	}

	if c.UsageMissing {
		return nil, false, nil
	}

	return &client.UsageResponse{}, true, nil
}
