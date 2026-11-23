package mock_client

func (c *Client) hasTeam(identifier string) bool {
	for _, t := range c.teams {
		if t.Identifier == identifier {
			return true
		}
	}

	return false
}
