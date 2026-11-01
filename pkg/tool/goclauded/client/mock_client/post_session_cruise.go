package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) PostSessionCruise(_ context.Context, _ string, _ client.PostSessionCruiseJSONRequestBody) (*client.CruiseResponse, error) {
	return &client.CruiseResponse{}, c.Failure
}
