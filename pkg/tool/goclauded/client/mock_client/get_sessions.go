package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetSessions(_ context.Context, _ *client.GetSessionsParams) (*client.SessionListResponse, error) {
	return &client.SessionListResponse{}, c.Failure
}
