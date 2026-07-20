package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) Setting() (*gitlab.Settings, error) {
	result, _, e := c.client.Settings.GetSettings()

	return result, e
}
