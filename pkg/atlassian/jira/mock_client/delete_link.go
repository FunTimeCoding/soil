package mock_client

func (c *Client) DeleteLink(_ string) error {
	if c.Failure != nil {
		return c.Failure
	}

	return c.failures.Caller()
}
