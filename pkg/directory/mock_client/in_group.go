package mock_client

func (c *Client) InGroup(account string) (bool, error) {
	if c.Failure != nil {
		return false, c.Failure
	}

	return c.members[account], nil
}
