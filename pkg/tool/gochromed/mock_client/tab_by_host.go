package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chromium/tab"
	"strings"
)

func (c *Client) TabByHost(s string) *tab.Tab {
	for _, t := range c.tabs {
		if strings.Contains(t.Locator, s) {
			return t
		}
	}

	return nil
}
