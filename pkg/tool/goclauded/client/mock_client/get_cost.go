package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetCost(_ context.Context, _ *client.GetCostParams) (*client.CostResponse, error) {
	return &client.CostResponse{}, c.Failure
}
