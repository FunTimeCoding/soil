package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
)

func (c *Client) PostGroupNameMember(
	_ context.Context,
	_ string,
	_ client.PostGroupNameMemberJSONRequestBody,
) (*client.Group, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return new(client.Group), nil
}
