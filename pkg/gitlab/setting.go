package gitlab

import "gitlab.com/gitlab-org/api/client-go/v3"

func (c *Client) Setting() (*gitlab.Settings, error) {
	result, _, e := c.client.Settings.GetSettings()

	return result, wrapError(e)
}
