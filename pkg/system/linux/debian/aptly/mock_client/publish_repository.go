package mock_client

func (c *Client) PublishRepository(
	_ string,
	_ string,
	_ []string,
	_ string,
) error {
	return c.fail
}
