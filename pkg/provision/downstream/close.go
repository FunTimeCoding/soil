package downstream

func (c *Client) Close() {
	c.pool.Close()
}
