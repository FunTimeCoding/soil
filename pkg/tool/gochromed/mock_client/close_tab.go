package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chromium/tab"
	"slices"
)

func (c *Client) CloseTab(identifier string) error {
	c.tabs = slices.DeleteFunc(
		c.tabs,
		func(t *tab.Tab) bool {
			return t.Identifier == identifier
		},
	)

	return nil
}
