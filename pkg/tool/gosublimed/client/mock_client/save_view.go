package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
)

func (c *Client) SaveView(
	_ context.Context,
	_ int,
	_ client.SaveViewJSONRequestBody,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.Failure
}
