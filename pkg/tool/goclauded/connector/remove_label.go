package connector

func (c *Client) RemoveLabel(
	sessionIdentifier string,
	key string,
	from string,
) error {
	return c.SetLabel(sessionIdentifier, key, "", from)
}
