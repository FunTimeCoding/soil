package mock_client

import (
	"github.com/funtimecoding/soil/pkg/browser/chromium/tab"
	"github.com/funtimecoding/soil/pkg/browser/constant"
)

func (c *Client) AddTab(
	identifier string,
	title string,
	locator string,
) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	t := tab.Tab{
		Identifier: identifier,
		Title:      title,
		Type:       constant.ChromiumPageTabType,
		Locator:    locator,
	}
	c.tabs = append(c.tabs, &t)
}
