package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (c *Client) PostContext(
	_ context.Context,
	b client.PostContextJSONRequestBody,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return c.Failure
	}

	c.ContextsAdded = append(c.ContextsAdded, b)

	return nil
}
