package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
)

func (c *Client) AppendTag(
	_ context.Context,
	_ int32,
	_ client.AppendTagJSONRequestBody,
) (*client.Link, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return new(client.Link), nil
}
