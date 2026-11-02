package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
)

func (c *Client) GetItem(
	_ context.Context,
	_ string,
) (*client.Item, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return new(client.Item), nil
}
