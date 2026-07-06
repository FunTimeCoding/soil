package github

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/github/workflow"
)

func (c *Client) MustWorkflows(
	owner string,
	name string,
) []*workflow.Workflow {
	result, e := c.Workflows(owner, name)
	errors.PanicOnError(e)

	return result
}
