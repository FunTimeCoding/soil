package band

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustScreenSettings() *ScreenSettingsResponse {
	result, e := c.ScreenSettings()
	errors.PanicOnError(e)

	return result
}
