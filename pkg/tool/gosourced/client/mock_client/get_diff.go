package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
)

func (c *Client) GetDiff(
	_ context.Context,
	p *client.GetDiffParams,
) (client.DiffResponse, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.Asked = append(c.Asked, p.Module)

	if c.Failure != nil {
		return client.DiffResponse{}, c.Failure
	}

	return c.Diff, nil
}
