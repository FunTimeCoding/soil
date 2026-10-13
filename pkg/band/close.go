package band

func (c *Client) Close() {
	c.pool.Close()
}
