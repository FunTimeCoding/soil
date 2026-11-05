package mock_client

func (c *Client) DeleteInternet(_ int32) error {
	if c.Failure != nil {
		return c.Failure
	}

	return nil
}
