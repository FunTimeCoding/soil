package mock_client

func (c *Client) Reveal(_ string) (string, bool) {
	return "", false
}
