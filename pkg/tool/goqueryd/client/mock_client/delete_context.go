package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (c *Client) DeleteContext(
	_ context.Context,
	p *client.DeleteContextParams,
) (bool, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return false, c.Failure
	}

	c.ContextsDeleted = append(c.ContextsDeleted, *p)

	return false, nil
}
