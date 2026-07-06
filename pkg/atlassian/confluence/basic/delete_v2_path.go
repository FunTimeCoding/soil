package basic

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (c *Client) DeleteV2Path(p string) (string, error) {
	return c.DeleteV2(
		locator.New(c.host).Base(constant.Base).Path(p).String(),
	)
}
