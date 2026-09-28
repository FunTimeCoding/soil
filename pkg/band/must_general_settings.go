package band

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustGeneralSettings() *GeneralSettingsResponse {
	result, e := c.GeneralSettings()
	errors.PanicOnError(e)

	return result
}
