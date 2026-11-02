package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
)

func (c *Client) OpenFile(
	_ context.Context,
	_ client.OpenFileJSONRequestBody,
) (*client.View, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return new(client.View), nil
}
