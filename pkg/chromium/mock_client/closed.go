package mock_client

func (c *Client) Closed() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.closed
}
