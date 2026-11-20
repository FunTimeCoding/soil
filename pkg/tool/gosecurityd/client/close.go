package client

func (c *Client) Close() {
	c.pool.Close()
}
