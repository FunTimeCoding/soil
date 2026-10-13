package openai

func (c *Client) Close() {
	c.pool.Close()
}
