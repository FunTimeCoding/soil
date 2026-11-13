package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
)

func (c *Client) PostCheck(
	_ context.Context,
	b client.CheckRequest,
) (client.CheckResponse, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.Asked = append(c.Asked, b.Root)

	if c.Failure != nil {
		return client.CheckResponse{}, c.Failure
	}

	return c.Check, nil
}
