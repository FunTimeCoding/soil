package fosite_client

func (c *Client) IsPublic() bool {
	return c.Row.Public
}
