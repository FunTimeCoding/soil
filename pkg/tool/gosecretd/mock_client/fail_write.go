package mock_client

func (c *Client) FailWrite(e error) {
	c.failWrite = e
}
