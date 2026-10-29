package mock_client

func (c *Client) SetMember(
	account string,
	member bool,
) {
	c.members[account] = member
}
