package anthropic

func (c *Client) Close() {
	c.pool.Close()
}
