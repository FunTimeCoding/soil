package pushover

func (c *Client) Close() {
	c.pool.Close()
}
