package mock_client

import "github.com/funtimecoding/soil/pkg/github/workflow"

func (c *Client) Workflows(_ string, _ string) ([]*workflow.Workflow, error) {
	return nil, c.Failure
}
