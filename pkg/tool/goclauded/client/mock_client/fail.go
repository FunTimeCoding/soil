package mock_client

func (c *Client) Fail(e error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.Failure = e
}
