package mock_client

import "slices"

func (c *Client) Writes() []string {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return slices.Clone(c.writes)
}
