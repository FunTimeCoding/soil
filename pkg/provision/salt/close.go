package salt

func (c *Client) Close() {
	c.basic.Close()
}
