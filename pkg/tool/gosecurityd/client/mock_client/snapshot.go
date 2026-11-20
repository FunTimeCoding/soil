package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"

func (c *Client) Snapshot(name string) (client.SystemSnapshot, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	result, okay := c.snapshots[name]

	return result, okay
}
