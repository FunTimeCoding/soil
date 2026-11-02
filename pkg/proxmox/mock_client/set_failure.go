package mock_client

func (c *Client) SetFailure(e error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.failure = e
}
