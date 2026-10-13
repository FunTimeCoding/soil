package alertmanager

func (c *Client) Close() {
	c.pool.Close()
}
