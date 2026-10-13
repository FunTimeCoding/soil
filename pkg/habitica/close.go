package habitica

func (c *Client) Close() {
	c.basic.Close()
}
