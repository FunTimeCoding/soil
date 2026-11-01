package mock_client

func (c *Client) DeleteTag(_ int64, _ string) error {
	return c.failure
}
