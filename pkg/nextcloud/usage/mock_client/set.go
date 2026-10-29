package mock_client

import "github.com/funtimecoding/soil/pkg/nextcloud/usage"

func (c *Client) Set(
	files int64,
	shares int64,
) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.usage = usage.New(files, shares)
}
