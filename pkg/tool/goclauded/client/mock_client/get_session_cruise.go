package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetSessionCruise(_ context.Context, _ string) (*client.CruiseResponse, error) {
	return &client.CruiseResponse{}, c.Failure
}
