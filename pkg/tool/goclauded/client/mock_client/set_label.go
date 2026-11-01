package mock_client

func (c *Client) SetLabel(_ string, _ string, _ string, _ string) error {
	return c.Failure
}
