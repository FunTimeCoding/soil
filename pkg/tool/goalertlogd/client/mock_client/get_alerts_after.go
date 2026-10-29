package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
)

func (c *Client) GetAlertsAfter(
	_ context.Context,
	_ *client.GetAlertsAfterParams,
) ([]client.AlertsResponse, error) {
	return nil, c.Failure
}
