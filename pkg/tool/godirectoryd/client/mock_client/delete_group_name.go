package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
)

func (c *Client) DeleteGroupName(
	_ context.Context,
	_ string,
) (*client.Receipt, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return new(client.Receipt), nil
}
