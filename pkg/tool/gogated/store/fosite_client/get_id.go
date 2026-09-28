package fosite_client

func (c *Client) GetID() string {
	return c.Row.Identifier
}
