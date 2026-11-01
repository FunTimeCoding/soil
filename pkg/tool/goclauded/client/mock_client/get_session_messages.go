package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetSessionMessages(_ context.Context, _ string) (*client.MessagesResponse, error) {
	return &client.MessagesResponse{}, c.Failure
}
