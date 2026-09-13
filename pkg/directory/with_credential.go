package directory

func (c *Client) WithCredential(
	bind string,
	password string,
) *Client {
	c.bind = bind
	c.password = password

	return c
}
