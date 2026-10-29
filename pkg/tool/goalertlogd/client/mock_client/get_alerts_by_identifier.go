package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
)

func (c *Client) GetAlertsByIdentifier(
	_ context.Context,
	_ *client.GetAlertsByIdentifierParams,
) ([]client.AlertsResponse, error) {
	return nil, c.Failure
}
