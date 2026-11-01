package mock_client

func (c *Client) Trace(
	_ int64,
	job int64,
) (string, error) {
	return c.traces[job], nil
}
