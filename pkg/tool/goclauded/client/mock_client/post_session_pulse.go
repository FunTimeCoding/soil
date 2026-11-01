package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) PostSessionPulse(_ context.Context, _ string, _ client.PostSessionPulseJSONRequestBody) (*client.DeliveryResponse, error) {
	return &client.DeliveryResponse{}, c.Failure
}
