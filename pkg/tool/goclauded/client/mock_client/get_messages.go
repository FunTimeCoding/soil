package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetMessages(_ context.Context, _ *client.GetMessagesParams) (*client.MessageReadResponse, error) {
	return &client.MessageReadResponse{}, c.Failure
}
