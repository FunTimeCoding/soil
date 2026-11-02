package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/client"
)

func (c *Client) PostEvent(
	_ context.Context,
	_ client.PostEventJSONRequestBody,
) (*client.EventResponse, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return new(client.EventResponse), nil
}
