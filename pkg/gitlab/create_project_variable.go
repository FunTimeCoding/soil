package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) CreateProjectVariable(
	project int64,
	key string,
	value string,
) (*gitlab.ProjectVariable, error) {
	result, _, e := c.client.ProjectVariables.CreateVariable(
		project,
		&gitlab.CreateProjectVariableOptions{
			Key:   &key,
			Value: &value,
			Raw:   new(true),
		},
	)

	return result, e
}
