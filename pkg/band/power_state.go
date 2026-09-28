package band

import "github.com/funtimecoding/soil/pkg/band/constant"

func (c *Client) PowerState() (*PowerStateResponse, error) {
	return enumerate[PowerStateResponse](c, constant.AssociatedPowerResource)
}
