package basic

import (
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
)

func (c *Client) Propfind() error {
	_, e := c.requester.Bytes(request.New(webConstant.Propfind, ""))

	return e
}
