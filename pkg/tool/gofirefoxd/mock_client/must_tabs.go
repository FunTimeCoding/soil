package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/firefox/tab"
)

func (c *Client) MustTabs() []*tab.Tab {
	result, e := c.Tabs()
	errors.PanicOnError(e)

	return result
}
