package chromium

import (
	"github.com/funtimecoding/soil/pkg/chromium/tab"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustTabByHost(s string) *tab.Tab {
	result, e := c.TabByHost(s)
	errors.PanicOnError(e)

	return result
}
