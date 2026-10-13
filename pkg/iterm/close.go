package iterm

func (c *Client) Close() {
	c.pool.Close()
}
