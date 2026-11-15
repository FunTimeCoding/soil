package mock_client

func (c *Client) Transition(
	string,
	string,
) error {
	if c.Failure != nil {
		return c.Failure
	}

	return c.failures.Caller()
}
