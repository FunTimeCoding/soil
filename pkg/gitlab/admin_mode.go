package gitlab

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) AdminMode(
	x context.Context,
	on bool,
) (*gitlab.Settings, error) {
	result, _, e := c.client.Settings.UpdateSettings(
		&gitlab.UpdateSettingsOptions{AdminMode: new(on)},
		gitlab.WithContext(x),
	)

	return result, wrapError(e)
}
