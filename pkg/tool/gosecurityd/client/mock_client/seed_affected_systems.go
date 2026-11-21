package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"

func (c *Client) SeedAffectedSystems(v *client.AffectedSystems) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.affected[v.Identifier] = v
}
