package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"

func (c *Client) SeedScoreStatus(v client.ScoreStatus) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.scoreStatus = v
}
