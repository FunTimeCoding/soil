package band

import (
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/band/response"
)

func (c *Client) ScreenService() (*response.ScreenService, error) {
	return enumerate[response.ScreenService](c, constant.ScreenServiceResource)
}
