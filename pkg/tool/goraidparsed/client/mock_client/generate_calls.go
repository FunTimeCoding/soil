package mock_client

func (c *Client) GenerateCalls() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.Calls
}
