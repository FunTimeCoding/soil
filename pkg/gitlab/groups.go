package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) Groups() ([]*gitlab.Group, error) {
	result, _, e := c.client.Groups.ListGroups(nil, nil)

	return result, wrapError(e)
}
