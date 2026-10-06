package band

import (
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/band/response"
)

func (c *Client) Consent() (*response.Consent, error) {
	return get[response.Consent](c, constant.ConsentResource)
}
