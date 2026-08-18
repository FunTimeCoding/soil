package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"

func (c *Client) Statistics() *client.Statistics {
	return c.Stats
}
