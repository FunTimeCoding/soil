package author

func (c *Client) Close() {
	c.pool.Close()
}
