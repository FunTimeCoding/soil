package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetSessionsFind(_ context.Context, _ *client.GetSessionsFindParams) (*client.FindResponse, error) {
	return &client.FindResponse{}, c.Failure
}
