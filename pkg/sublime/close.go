package sublime

func (c *Client) Close() {
	c.pool.Close()
}
