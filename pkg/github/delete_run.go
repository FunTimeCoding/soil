package github

import "github.com/funtimecoding/soil/pkg/github/refusal"

func (c *Client) DeleteRun(
	owner string,
	name string,
	run int64,
) error {
	_, e := c.client.Actions.DeleteWorkflowRun(c.context, owner, name, run)

	return refusal.Classify(e)
}
