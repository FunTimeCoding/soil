package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
)

func (c *Client) GetStatus(_ context.Context) (*client.StatusResponse, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return &client.StatusResponse{}, nil
}
