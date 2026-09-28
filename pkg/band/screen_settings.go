package band

import "github.com/funtimecoding/soil/pkg/band/constant"

func (c *Client) ScreenSettings() (*ScreenSettingsResponse, error) {
	return get[ScreenSettingsResponse](c, constant.ScreenSettingsResource)
}
