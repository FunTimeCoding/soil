package postgres

func (c *Client) Close() {
	c.pool.Close()
}
