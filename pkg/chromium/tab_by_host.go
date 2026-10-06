package chromium

import (
	"github.com/funtimecoding/soil/pkg/chromium/tab"
	"strings"
)

func (c *Client) TabByHost(s string) (*tab.Tab, error) {
	tabs, e := c.Tabs()

	if e != nil {
		return nil, e
	}

	for _, t := range tabs {
		if strings.Contains(t.Locator, s) {
			return t, nil
		}
	}

	return nil, nil
}
