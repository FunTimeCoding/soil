package mock_client

func (c *Client) DeletePackageVersion(_ string, _ int64) error {
	return c.Failure
}
