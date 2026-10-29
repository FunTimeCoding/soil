package closing_client

func (c *Client) Close() {
	c.closed++
}
