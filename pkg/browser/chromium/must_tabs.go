package chromium

import (
	"github.com/funtimecoding/soil/pkg/browser/chromium/tab"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustTabs() []*tab.Tab {
	result, e := c.Tabs()
	errors.PanicOnError(e)

	return result
}
