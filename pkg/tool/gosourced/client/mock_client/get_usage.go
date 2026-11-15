package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
)

func (c *Client) GetUsage(
	_ context.Context,
	p *client.GetUsageParams,
) (client.Usage, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.Asked = append(c.Asked, p.Module)

	if c.Failure != nil {
		return client.Usage{}, c.Failure
	}

	return c.Usage, nil
}
