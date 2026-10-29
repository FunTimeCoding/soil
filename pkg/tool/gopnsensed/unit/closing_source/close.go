package closing_source

func (c *Client) Close() {
	c.closed++
}
