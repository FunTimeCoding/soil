package client

func (c *Client) endSessionLocator() string {
	c.ensureProvider()
	var m discoveryMetadata

	if e := c.provider.Claims(&m); e != nil {
		return ""
	}

	return m.EndSessionLocator
}
