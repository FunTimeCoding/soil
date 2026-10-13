package ollama

func (c *Client) Close() {
	c.pool.Close()
}
