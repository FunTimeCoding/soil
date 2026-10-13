package confluence

func (c *Client) Close() {
	c.basic.Close()
}
