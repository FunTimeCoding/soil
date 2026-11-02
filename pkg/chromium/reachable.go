package chromium

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
)

func (c *Client) Reachable(x context.Context) error {
	_, e := c.requester.Bytes(request.Get(constant.VersionPath).WithContext(x))

	return e
}
