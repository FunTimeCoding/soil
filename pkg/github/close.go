package github

func (c *Client) Close() {
	c.pool.Close()
}
