package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (c *Client) PostRechunk(context.Context) (*client.RechunkResult, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return &client.RechunkResult{}, nil
}
