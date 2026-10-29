package mock_client

func (c *Client) SilenceExists(_ string) (bool, error) {
	return false, c.failure
}
