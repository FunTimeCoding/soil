package mock_client

func (c *Client) Create(
	_ string,
	_ string,
	_ string,
	_ string,
	_ string,
) error {
	return c.failure
}
