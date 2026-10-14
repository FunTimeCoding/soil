package assistant

func (c *Client) Close() {
	c.connection.Stop()
}
