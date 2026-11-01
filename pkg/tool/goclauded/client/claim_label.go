package client

func (c *Client) ClaimLabel(
	sessionIdentifier string,
	key string,
	value string,
	from string,
) (bool, error) {
	return c.swapLabel(sessionIdentifier, key, "", value, from)
}
