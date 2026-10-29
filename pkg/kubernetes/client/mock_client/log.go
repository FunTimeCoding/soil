package mock_client

func (c *Client) Log(_ string, _ string, _ string) (string, error) {
	return "", c.Failure
}
