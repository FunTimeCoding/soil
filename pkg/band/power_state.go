package band

import (
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/band/response"
)

func (c *Client) PowerState() (*response.PowerState, error) {
	return enumerate[response.PowerState](c, constant.AssociatedPowerResource)
}
