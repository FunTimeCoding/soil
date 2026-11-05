package mock_client

func (c *Client) Highstates() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.highstates
}
