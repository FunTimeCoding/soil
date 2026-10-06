package band

import (
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/band/response"
)

func (c *Client) ScreenSettings() (*response.ScreenSettings, error) {
	return get[response.ScreenSettings](c, constant.ScreenSettingsResource)
}
