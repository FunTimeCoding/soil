package telemetry

func (c *Client) Close() {
	c.pool.Close()
}
