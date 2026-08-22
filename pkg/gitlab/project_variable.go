package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) ProjectVariable(
	project int64,
	key string,
) (*gitlab.ProjectVariable, error) {
	result, _, e := c.client.ProjectVariables.GetVariable(project, key, nil)

	return result, wrapError(e)
}
