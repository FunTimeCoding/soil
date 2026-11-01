package mock_client

func (c *Client) SeedTrace(
	job int64,
	trace string,
) {
	c.traces[job] = trace
}
