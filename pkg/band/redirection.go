package band

import (
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/band/response"
)

func (c *Client) Redirection() (*response.Redirection, error) {
	return get[response.Redirection](c, constant.RedirectionResource)
}
