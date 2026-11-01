package mock_client

func (c *Client) DeleteRun(_ string, _ string, _ int64) error {
	return c.Failure
}
