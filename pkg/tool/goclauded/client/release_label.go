package client

func (c *Client) ReleaseLabel(
	sessionIdentifier string,
	key string,
	value string,
	from string,
) (bool, error) {
	return c.swapLabel(sessionIdentifier, key, value, "", from)
}
