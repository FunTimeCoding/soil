package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
)

func (c *Client) PostUser(
	_ context.Context,
	_ client.PostUserJSONRequestBody,
) (*client.User, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return new(client.User), nil
}
