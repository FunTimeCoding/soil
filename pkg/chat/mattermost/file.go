package mattermost

func (c *Client) File(identifier string) ([]byte, error) {
	result, r, e := c.client.GetFile(c.context, identifier)

	return result, wrapError(r, e)
}
