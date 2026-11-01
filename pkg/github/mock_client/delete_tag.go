package mock_client

func (c *Client) DeleteTag(_ string, _ string, _ string) error {
	return c.Failure
}
