package mock_client

func (c *Client) DeleteJob(_ string, _ string) error {
	return c.Failure
}
