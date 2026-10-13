package technitium

func (c *Client) Close() {
	c.basic.Close()
}
