package raid

func (c *Client) Close() {
	c.pool.Close()
}
