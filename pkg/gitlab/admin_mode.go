package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) AdminMode(on bool) (*gitlab.Settings, error) {
	// AdminMode is persisted
	result, _, e := c.client.Settings.UpdateSettings(
		&gitlab.UpdateSettingsOptions{AdminMode: new(on)},
	)

	return result, wrapError(e)
}
