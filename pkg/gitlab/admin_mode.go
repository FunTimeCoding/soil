package gitlab

import "gitlab.com/gitlab-org/api/client-go/v3"

func (c *Client) AdminMode(on bool) (*gitlab.Settings, error) {
	result, _, e := c.client.Settings.UpdateSettings(
		&gitlab.UpdateSettingsOptions{AdminMode: new(on)},
	)

	return result, wrapError(e)
}
