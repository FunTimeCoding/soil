package kestra

func (c *Client) Close() {
	c.pool.Close()
}
