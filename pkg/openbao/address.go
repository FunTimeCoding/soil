package openbao

func (c *Client) Address() string {
	return c.client.Address()
}
