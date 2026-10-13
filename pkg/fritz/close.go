package fritz

func (c *Client) Close() {
	c.pool.Close()
}
