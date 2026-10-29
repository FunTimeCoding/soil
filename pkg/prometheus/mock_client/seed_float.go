package mock_client

func (c *Client) SeedFloat(
	q string,
	value float64,
) {
	c.floats[q] = value
}
