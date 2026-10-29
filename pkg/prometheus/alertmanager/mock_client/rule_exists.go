package mock_client

func (c *Client) RuleExists(_ string) (bool, error) {
	return false, c.failure
}
