package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetSessionPeek(_ context.Context, _ string) (*client.PeekResponse, error) {
	return &client.PeekResponse{}, c.Failure
}
