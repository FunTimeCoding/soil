package mock_client

func (c *Client) DeleteInterface(_ int32) error {
	return c.Failure
}
