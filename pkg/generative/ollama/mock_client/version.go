package mock_client

func (c *Client) Version() (string, error) {
	if c.Failure != nil {
		return "", c.Failure
	}

	return c.Release, nil
}
