package mock_client

func (c *Client) Configuration() (map[string]string, error) {
	return nil, c.Failure
}
