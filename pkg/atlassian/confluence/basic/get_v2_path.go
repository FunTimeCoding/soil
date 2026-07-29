package basic

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (c *Client) GetV2Path(p string) (string, error) {
	return c.GetV2(
		locator.New(c.host).Base(constant.ConfluenceBase).Path(p).String(),
	)
}
