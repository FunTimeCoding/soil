package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"

func (c *Client) SeedFindingGroup(v *client.FindingGroup) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.groups = append(c.groups, v)
}
