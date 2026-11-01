package mock_client

func (c *Client) DeletePackage(_ int64, _ int64) error {
	return c.failure
}
