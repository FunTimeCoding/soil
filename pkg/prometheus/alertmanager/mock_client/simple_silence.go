package mock_client

func (c *Client) SimpleSilence(_ string) (string, error) {
	return "", c.failure
}
