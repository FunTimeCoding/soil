package failing_jira

func (c *Client) Transition(
	string,
	string,
) error {
	return c.failure
}
