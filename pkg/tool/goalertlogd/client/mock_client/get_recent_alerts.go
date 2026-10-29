package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
)

func (c *Client) GetRecentAlerts(
	_ context.Context,
	_ *client.GetRecentAlertsParams,
) ([]client.AlertsResponse, error) {
	return nil, c.Failure
}
