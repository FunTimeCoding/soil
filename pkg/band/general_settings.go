package band

import (
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/band/response"
)

func (c *Client) GeneralSettings() (*response.GeneralSettings, error) {
	return get[response.GeneralSettings](c, constant.GeneralSettingsResource)
}
