package connector

func (c *Client) Close() {
	c.pool.Close()
}
