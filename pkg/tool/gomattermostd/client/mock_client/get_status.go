package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"
)

func (c *Client) GetStatus(_ context.Context) (*client.Status, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return &client.Status{}, nil
}
