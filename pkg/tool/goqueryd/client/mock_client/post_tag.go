package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (c *Client) PostTag(
	_ context.Context,
	_ client.PostTagJSONRequestBody,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.Failure
}
