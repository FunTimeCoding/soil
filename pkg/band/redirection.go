package band

import "github.com/funtimecoding/soil/pkg/band/constant"

func (c *Client) Redirection() (*RedirectionResponse, error) {
	return get[RedirectionResponse](c, constant.RedirectionResource)
}
