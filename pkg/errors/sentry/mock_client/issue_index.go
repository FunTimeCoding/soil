package mock_client

func (c *Client) issueIndex(identifier string) int {
	for i, issue := range c.issues {
		if issue.Identifier == identifier {
			return i
		}
	}

	return -1
}
