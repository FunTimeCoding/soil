package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
)

func (c *Client) PostUserAccountPassword(
	_ context.Context,
	_ string,
	_ client.PostUserAccountPasswordJSONRequestBody,
) (*client.Receipt, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return new(client.Receipt), nil
}
