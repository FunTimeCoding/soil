package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (c *Client) PostTag(
	_ context.Context,
	b client.PostTagJSONRequestBody,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return c.Failure
	}

	c.Tagged = append(c.Tagged, b)

	return nil
}
