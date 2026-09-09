package gitlab

import "github.com/funtimecoding/soil/pkg/gitlab/variable"

func (c *Client) ProjectVariable(
	project int64,
	key string,
) (*variable.Variable, error) {
	result, _, e := c.client.ProjectVariables.GetVariable(project, key, nil)

	if e != nil {
		return nil, wrapError(e)
	}

	return variable.New(result), nil
}
