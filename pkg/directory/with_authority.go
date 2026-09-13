package directory

func (c *Client) WithAuthority(path string) *Client {
	c.authority = path

	return c
}
