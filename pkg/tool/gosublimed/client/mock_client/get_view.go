package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
)

func (c *Client) GetView(
	_ context.Context,
	_ int,
) (*client.View, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return new(client.View), nil
}
