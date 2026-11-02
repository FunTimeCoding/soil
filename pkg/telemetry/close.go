package telemetry

func (c *Client) Close() {
	c.client.Close()
}
