package github

import "github.com/funtimecoding/soil/pkg/github/workflow"

func (c *Client) Workflows(
	owner string,
	name string,
) ([]*workflow.Workflow, error) {
	result, _, e := c.client.Actions.ListWorkflows(
		c.context,
		owner,
		name,
		nil,
	)

	if e != nil {
		return nil, e
	}

	return workflow.NewSlice(result.Workflows), nil
}
