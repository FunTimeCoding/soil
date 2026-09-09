package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/variable"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) UpdateProjectVariable(
	project int64,
	key string,
	value string,
	protected bool,
	masked bool,
	literal bool,
) (*variable.Variable, error) {
	result, _, e := c.client.ProjectVariables.UpdateVariable(
		project,
		key,
		&gitlab.UpdateProjectVariableOptions{
			Value:     &value,
			Protected: &protected,
			Masked:    &masked,
			Raw:       &literal,
		},
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return variable.New(result), nil
}
