package hetzner

func (c *Client) Close() {
	c.pool.Close()
}
