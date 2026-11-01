package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetSessionDetail(_ context.Context, _ string) (*client.SessionDetailResponse, error) {
	return &client.SessionDetailResponse{}, c.Failure
}
