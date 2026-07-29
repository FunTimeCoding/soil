package basic

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (c *Client) GetPath(path string) (string, error) {
	return c.Get(
		locator.New(c.host).Base(constant.ConfluenceOldBase).Path(path).String(),
	)
}
