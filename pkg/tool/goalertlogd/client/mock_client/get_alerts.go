package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
)

func (c *Client) GetAlerts(
	_ context.Context,
	_ *client.GetAlertsParams,
) ([]client.AlertsResponse, error) {
	return nil, c.Failure
}
