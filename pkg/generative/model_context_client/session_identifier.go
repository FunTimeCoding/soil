package model_context_client

func (c *Client) SessionIdentifier() string {
	return c.session.ID()
}
