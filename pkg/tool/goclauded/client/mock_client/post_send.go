package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) PostSend(_ context.Context, _ client.PostSendJSONRequestBody) (*client.DeliveryResponse, error) {
	return &client.DeliveryResponse{}, c.Failure
}
