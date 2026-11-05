package mock_client

func (c *Client) DeviceTagNames(_ string) ([]string, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
