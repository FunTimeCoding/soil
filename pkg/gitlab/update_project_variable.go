package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) UpdateProjectVariable(
	project int64,
	key string,
	value string,
) (*gitlab.ProjectVariable, error) {
	result, _, e := c.client.ProjectVariables.UpdateVariable(
		project,
		key,
		&gitlab.UpdateProjectVariableOptions{Value: &value},
	)

	return result, wrapError(e)
}
