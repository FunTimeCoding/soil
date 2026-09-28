package band

import "github.com/funtimecoding/soil/pkg/band/constant"

func (c *Client) Consent() (*ConsentResponse, error) {
	return get[ConsentResponse](c, constant.ConsentResource)
}
