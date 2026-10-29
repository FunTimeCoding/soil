package mock_client

func (c *Client) DeleteSilence(_ string) error {
	return c.failure
}
