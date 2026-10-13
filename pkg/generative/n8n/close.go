package n8n

func (c *Client) Close() {
	c.pool.Close()
}
