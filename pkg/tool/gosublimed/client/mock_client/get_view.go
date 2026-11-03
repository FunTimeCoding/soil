package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
)

func (c *Client) GetView(
	_ context.Context,
	identifier int,
) (*client.View, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	i, e := c.index(identifier)

	if e != nil {
		return nil, e
	}

	result := c.ViewList[i]

	return &result, nil
}
