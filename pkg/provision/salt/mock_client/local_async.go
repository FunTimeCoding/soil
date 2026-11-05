package mock_client

func (c *Client) LocalAsync(
	_ string,
	_ string,
	_ []string,
) (string, error) {
	if e := c.failures.Caller(); e != nil {
		return "", e
	}

	return "", nil
}
