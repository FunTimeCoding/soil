package mock_client

import "slices"

func (c *Client) Requested() []int {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return slices.Clone(c.requested)
}
