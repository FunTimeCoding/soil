package mock_client

func (c *Client) Generations() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return len(c.Generated)
}
