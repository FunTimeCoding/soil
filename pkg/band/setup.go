package band

import (
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/band/response"
)

func (c *Client) Setup() (*response.Setup, error) {
	return get[response.Setup](c, constant.SetupResource)
}
