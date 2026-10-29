package mock_client

func (c *Client) index(identifier string) int {
	for i, h := range c.HostList {
		if h.Identifier == identifier {
			return i
		}
	}

	return -1
}
