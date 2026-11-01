package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
)

func (c *Client) GetDiff(
	_ context.Context,
	module string,
	_ string,
	_ string,
) (client.DiffResponse, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.Asked = append(c.Asked, module)

	if c.Failure != nil {
		return client.DiffResponse{}, c.Failure
	}

	return c.Diff, nil
}
