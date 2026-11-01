package mock_client

func (c *Client) DeleteImage(_ int64, _ int64, _ string) error {
	return c.failure
}
