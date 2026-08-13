package github

func (c *Client) DeleteRun(
	owner string,
	name string,
	run int64,
) error {
	_, e := c.client.Actions.DeleteWorkflowRun(c.context, owner, name, run)

	return e
}
