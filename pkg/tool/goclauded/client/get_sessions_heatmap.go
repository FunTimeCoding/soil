package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetSessionsHeatmap(
	x context.Context,
	params *client.GetSessionsHeatmapParams,
) (*client.HeatmapResponse, error) {
	r, e := c.generated.GetSessionsHeatmapWithResponse(x, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get sessions heatmap", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
