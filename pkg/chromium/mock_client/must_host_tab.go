package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chromium/tab"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustHostTab(s string) *tab.Tab {
	result, e := c.HostTab(s)
	errors.PanicOnError(e)

	return result
}
