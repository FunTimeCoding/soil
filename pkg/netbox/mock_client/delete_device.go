package mock_client

func (c *Client) DeleteDevice(_ int32) error {
	return c.Failure
}
