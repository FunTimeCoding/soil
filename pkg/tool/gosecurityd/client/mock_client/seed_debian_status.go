package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"

func (c *Client) SeedDebianStatus(v *client.DebianStatus) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.statuses[v.Identifier] = v
}
