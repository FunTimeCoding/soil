package mock_client

func (c *Client) hasChannel(
	team string,
	identifier string,
) bool {
	for _, h := range c.channels[team] {
		if h.Identifier == identifier {
			return true
		}
	}

	return false
}
