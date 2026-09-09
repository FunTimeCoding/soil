package mock_client

func (c *Client) ResolveProject(_ string) (int64, error) {
	return 0, nil
}
