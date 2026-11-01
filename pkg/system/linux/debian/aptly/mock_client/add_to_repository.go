package mock_client

func (c *Client) AddToRepository(
	_ string,
	_ string,
	_ bool,
) error {
	return c.fail
}
