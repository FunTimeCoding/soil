package mock_client

func (c *Client) Calls() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.calls
}
