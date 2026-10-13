package opnsense

func (c *Client) Close() {
	c.basic.Close()
}
