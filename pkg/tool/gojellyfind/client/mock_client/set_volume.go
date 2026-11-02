package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
)

func (c *Client) SetVolume(
	_ context.Context,
	_ string,
	_ client.SetVolumeJSONRequestBody,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.Failure
}
