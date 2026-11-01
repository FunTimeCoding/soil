package mock_client

func (c *Client) Upload(
	_ string,
	_ string,
) error {
	return c.fail
}
