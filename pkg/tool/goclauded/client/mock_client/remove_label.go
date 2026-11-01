package mock_client

func (c *Client) RemoveLabel(_ string, _ string, _ string) error {
	return c.Failure
}
