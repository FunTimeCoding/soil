package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/firefox/tab"
)

func (c *Client) MustCreateTab(l string) *tab.Tab {
	result, e := c.CreateTab(l)
	errors.PanicOnError(e)

	return result
}
