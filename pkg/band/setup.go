package band

import "github.com/funtimecoding/soil/pkg/band/constant"

func (c *Client) Setup() (*SetupResponse, error) {
	return get[SetupResponse](c, constant.SetupResource)
}
