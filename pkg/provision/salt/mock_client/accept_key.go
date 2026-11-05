package mock_client

func (c *Client) AcceptKey(minion string) ([]string, error) {
	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	return []string{minion}, nil
}
