package mock_client

func (c *Client) ExecuteStandard(_ string, _ string, _ string, _ ...string) error {
	return c.Failure
}
