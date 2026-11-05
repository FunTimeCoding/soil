package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetCost(
	_ context.Context,
	p *client.GetCostParams,
) (*client.CostResponse, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	result := &client.CostResponse{Days: 1}

	if p != nil && p.Days != nil {
		result.Days = *p.Days
	}

	return result, nil
}
