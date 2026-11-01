package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetSessionLabel(_ context.Context, _ string) (*client.SessionLabelResponse, error) {
	return &client.SessionLabelResponse{}, c.Failure
}
