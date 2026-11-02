package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/client"
)

func (c *Client) GetEvents(
	_ context.Context,
	_ *client.GetEventsParams,
) ([]client.EventEntry, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
