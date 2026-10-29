package mock_client

func (c *Client) DeleteJobWatch(_ string, _ string) error {
	return c.Failure
}
