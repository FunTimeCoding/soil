package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"

func (c *Client) SeedUpgrade(v client.UpgradeGroup) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.upgrades = append(c.upgrades, v)
}
