package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
)

func (c *Client) GetGroup(_ context.Context) ([]*client.Group, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
