package mock_client

func (c *Client) DeleteComment(
	_ string,
	_ string,
) error {
	if c.Failure != nil {
		return c.Failure
	}

	return c.failures.Caller()
}
