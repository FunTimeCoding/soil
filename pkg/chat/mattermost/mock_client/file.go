package mock_client

func (c *Client) File(_ string) ([]byte, error) {
	return nil, c.Failure
}
