package prometheus

func (c *Client) Close() {
	c.pool.Close()
}
