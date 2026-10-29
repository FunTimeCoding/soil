package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
)

func (c *Client) GetTopAlerts(
	_ context.Context,
	_ *client.GetTopAlertsParams,
) ([]client.TopAlertsResponse, error) {
	return nil, c.Failure
}
