package directory

func (c *Client) WithFilter(
	user string,
	group string,
) *Client {
	c.userFilter = user
	c.groupFilter = group

	return c
}
