package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/tab"
)

func (c *Client) AddTab(
	identifier string,
	title string,
	locator string,
) {
	t := tab.Tab{
		Identifier: identifier,
		Title:      title,
		Type:       constant.PageTabType,
		Locator:    locator,
	}
	c.tabs = append(c.tabs, &t)
}
