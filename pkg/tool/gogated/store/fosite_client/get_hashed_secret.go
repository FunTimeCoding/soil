package fosite_client

func (c *Client) GetHashedSecret() []byte {
	return []byte(c.Row.Secret)
}
