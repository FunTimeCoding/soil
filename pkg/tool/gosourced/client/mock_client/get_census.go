package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
)

func (c *Client) GetCensus(
	_ context.Context,
	p *client.GetCensusParams,
) (client.Census, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.Asked = append(c.Asked, p.Module)

	if c.Failure != nil {
		return client.Census{}, c.Failure
	}

	return c.Census, nil
}
