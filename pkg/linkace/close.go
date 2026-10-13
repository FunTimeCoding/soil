package linkace

func (c *Client) Close() {
	c.basic.Close()
}
