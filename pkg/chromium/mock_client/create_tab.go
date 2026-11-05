package mock_client

func (c *Client) CreateTab(_ string) (string, error) {
	if c.Failure != nil {
		return "", c.Failure
	}

	return "", nil
}
