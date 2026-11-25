package mock_client

func (c *Client) SetApprovals(
	project int64,
	request int64,
	usernames ...string,
) {
	c.approvals[[2]int64{project, request}] = usernames
}
