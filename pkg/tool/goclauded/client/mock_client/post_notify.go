package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) PostNotify(_ context.Context, _ client.PostNotifyJSONRequestBody) (*client.DeliveryResponse, error) {
	return &client.DeliveryResponse{}, c.Failure
}
