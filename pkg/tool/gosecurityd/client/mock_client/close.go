package mock_client

func (c *Client) Close() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.Closed++
}
