package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
)

func (c *Client) PlaybackCommand(
	_ context.Context,
	_ string,
	_ client.PlaybackCommandJSONRequestBody,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.Failure
}
