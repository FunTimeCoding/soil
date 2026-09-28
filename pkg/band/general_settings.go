package band

import "github.com/funtimecoding/soil/pkg/band/constant"

func (c *Client) GeneralSettings() (*GeneralSettingsResponse, error) {
	return get[GeneralSettingsResponse](c, constant.GeneralSettingsResource)
}
