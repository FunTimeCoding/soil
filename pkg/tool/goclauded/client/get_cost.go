package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetCost(
	x context.Context,
	params *client.GetCostParams,
) (*client.CostResponse, error) {
	r, e := c.generated.GetCostWithResponse(x, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get cost", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
