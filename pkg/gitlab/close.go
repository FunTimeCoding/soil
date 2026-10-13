package gitlab

func (c *Client) Close() {
	c.pool.Close()
}
