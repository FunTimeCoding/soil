package mock_client

func (c *Client) SeedFailure(
	q string,
	e error,
) {
	c.failures[q] = e
}
