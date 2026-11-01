package mock_client

func (c *Client) UpdatePublish(
	_ string,
	_ string,
) error {
	return c.fail
}
