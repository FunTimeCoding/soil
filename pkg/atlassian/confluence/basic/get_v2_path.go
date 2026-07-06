package basic

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (c *Client) GetV2Path(p string) (string, error) {
	return c.GetV2(
		locator.New(c.host).Base(constant.Base).Path(p).String(),
	)
}
