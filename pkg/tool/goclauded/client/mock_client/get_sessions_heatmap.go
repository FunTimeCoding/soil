package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetSessionsHeatmap(_ context.Context, _ *client.GetSessionsHeatmapParams) (*client.HeatmapResponse, error) {
	return &client.HeatmapResponse{}, c.Failure
}
